package service

import (
	"context"
	"log/slog"

	"github.com/dijonron/recipe-box/pkg/grpc"
	pb "github.com/dijonron/recipe-box/proto/grpc/auth"
	auth "github.com/dijonron/recipe-box/services/auth/internal"

	g "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const (
	ERROR_OCCURED       string = "an unknown error occured"
	EMTPY_REQUEST       string = "request is nil"
	INVALID_CREDENTIALS string = "invalid credentials"
	MISSING_EMAIL       string = "email is required"
	MISSING_PASSWORD    string = "password is required"
)

var (
	errUnknownError       = status.Errorf(codes.Internal, ERROR_OCCURED)
	errEmptyRequest       = status.Errorf(codes.InvalidArgument, EMTPY_REQUEST)
	errInvalidCredentials = status.Error(codes.Unauthenticated, INVALID_CREDENTIALS)
	errMissingEmail       = status.Errorf(codes.InvalidArgument, MISSING_EMAIL)
	errMissingPassword    = status.Errorf(codes.InvalidArgument, MISSING_PASSWORD)
)

type authServer struct {
	auth auth.AuthManager
	pb.UnimplementedAuthServer
}

var _ pb.AuthServer = (*authServer)(nil)

func NewAuthServer(cfg grpc.ServerConfig, a auth.AuthManager) Server {
	if a == nil {
		slog.Error("auth manager is nil")
		return nil
	}

	var registers []grpc.Register
	h := &authServer{
		auth: a,
	}
	registers = append(registers, h)

	s := grpc.NewServer(cfg, registers)

	as := server{s}

	return as

}

func (s *authServer) Register(grpcServer *g.Server) {
	pb.RegisterAuthServer(grpcServer, s)
}

func (s *authServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := validateLoginRequest(req); err != nil {
		return &pb.LoginResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	token, err := s.auth.LoginUser(ctx, req.GetEmail(), req.GetPassword())
	if err != nil {
		slog.Info(INVALID_CREDENTIALS, "error", err)
		if err.Error() == auth.INVALID_CREDENTIALS {
			return &pb.LoginResponse{}, errInvalidCredentials
		}
		return &pb.LoginResponse{}, errUnknownError
	}

	return &pb.LoginResponse{
		Token:   token,
		Success: true,
	}, nil
}

func (s *authServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
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
