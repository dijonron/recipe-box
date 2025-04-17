package service

import (
	"context"
	"log/slog"
	"regexp"

	"github.com/dijonron/recipe-box/pkg/grpc"
	pb "github.com/dijonron/recipe-box/proto/grpc/user"
	u "github.com/dijonron/recipe-box/services/user/internal"

	g "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	ERROR_OCCURED         string = "an unknown error occured"
	EMTPY_REQUEST         string = "request is nil"
	FAILED_TO_CREATE_USER string = "failed to create user"
	INVALID_CREDENTIALS   string = "invalid credentials"
	INVALID_EMAIL         string = "invlaid email address"
	INVALID_PASSWORD      string = "invlaid password"
	MISSING_EMAIL         string = "email is required"
	MISSING_NAME          string = "name is required"
	MISSING_PASSWORD      string = "password is required"
)

var (
	errEmptyRequest       = status.Errorf(codes.InvalidArgument, EMTPY_REQUEST)
	errFailedToCreateUser = status.Errorf(codes.Internal, FAILED_TO_CREATE_USER)
	errInvalidCredentials = status.Error(codes.Unauthenticated, INVALID_CREDENTIALS)

	errInvalidEmail    = status.Errorf(codes.InvalidArgument, INVALID_EMAIL)
	errInvalidPassword = status.Errorf(codes.InvalidArgument, INVALID_PASSWORD)
	errMissingEmail    = status.Errorf(codes.InvalidArgument, MISSING_EMAIL)
	errMissingName     = status.Errorf(codes.InvalidArgument, MISSING_NAME)
	errMissingPassword = status.Errorf(codes.InvalidArgument, MISSING_PASSWORD)
	errUnknownError    = status.Errorf(codes.Internal, ERROR_OCCURED)
)

type userServer struct {
	user u.UserManager
	pb.UnimplementedUserServer
}

var _ pb.UserServer = (*userServer)(nil)

func NewUserServer(cfg grpc.ServerConfig, u u.UserManager) Server {
	if u == nil {
		slog.Error("user manager is nil")
		return nil
	}

	var registers []grpc.Register
	h := &userServer{
		user: u,
	}
	registers = append(registers, h)

	s := grpc.NewServer(cfg, registers)

	us := server{s}

	return us

}

func (s *userServer) Register(grpcServer *g.Server) {
	pb.RegisterUserServer(grpcServer, s)
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	name, email, password, err := validateCreateUserRequest(req)
	if err != nil {
		slog.Info("create user request failed", "error", err)
		return nil, err
	}

	user, err := s.user.CreateUser(ctx, name, email, password)
	if err != nil {
		// TODO: already exists error
		return nil, errFailedToCreateUser
	}

	userDetails := toUserDetails(user)
	resp := &pb.CreateUserResponse{
		User: userDetails,
	}
	return resp, nil
}

func (s *userServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := validateLoginRequest(req); err != nil {
		return &pb.LoginResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.user.Login(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		if err.Error() == u.INVALID_CREDENTIALS {
			return &pb.LoginResponse{}, errInvalidCredentials
		}
		slog.Debug("create user request failed", "error", err)
		return &pb.LoginResponse{}, errUnknownError
	}

	return &pb.LoginResponse{
		User: toUserDetails(user),
	}, nil
}

func (s *userServer) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, errMissingEmail
	}

	user, err := s.user.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found: %v", err)
	}

	resp := &pb.GetUserByEmailResponse{
		User: &pb.UserDetails{
			Name:         user.Name,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			TenantId:     user.TenantID,
			Role:         user.Role,
		},
	}
	return resp, nil

}

func (s *userServer) UpdateUserLastLogin(ctx context.Context, req *pb.UpdateUserLastLoginRequest) (*pb.UpdateUserLastLoginResponse, error) {
	email := req.GetEmail()
	if email == "" {
		return nil, errMissingEmail
	}

	err := s.user.UpdateUserLogin(ctx, email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user login: %v", err)
	}

	resp := &pb.UpdateUserLastLoginResponse{}
	return resp, nil
}

// validateCreateUserRequest validates the input for creating a new user.
func validateCreateUserRequest(req *pb.CreateUserRequest) (name, email, password string, error error) {
	if req == nil {
		return "", "", "", errEmptyRequest
	}

	name = req.GetName()
	if name == "" {
		return "", "", "", errMissingName
	}

	email = req.GetEmail()
	if email == "" {
		return "", "", "", errMissingEmail
	}
	if !isValidEmail(email) {
		return "", "", "", errInvalidEmail
	}

	password = req.GetPassword()
	if password == "" {
		return "", "", "", errMissingPassword
	}
	if !isValidPassword(password) {
		return "", "", "", errInvalidPassword
	}

	return name, email, password, nil
}

// isValidEmail validates whether the given email string is in a proper email format.
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// isValidPassword validates a password based on the following criteria:
func isValidPassword(password string) bool {
	// Password should be at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Must contain at least one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasDigit {
		return false
	}

	// Must contain at least one uppercase letter
	hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUppercase {
		return false
	}

	// Must contain at least one special character (e.g., !@#$%^&*)
	hasSpecialChar := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'",<>\./?\\|]`).MatchString(password)
	return hasSpecialChar
}

// validateLoginRequest validates the login request
func validateLoginRequest(req *pb.LoginRequest) error {
	if req == nil {
		return errEmptyRequest
	}

	if req.GetEmail() == "" {
		return errMissingEmail
	}
	if req.GetPassword() == "" {
		return errMissingPassword
	}
	return nil
}

func toUserDetails(user u.User) *pb.UserDetails {

	u := &pb.UserDetails{
		Id:           user.Id,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		TenantId:     user.TenantID,
		Role:         user.Role,
	}

	return u
}
