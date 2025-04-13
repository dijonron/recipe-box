package service

import (
	"log/slog"

	"github.com/dijonron/recipe-box/pkg/grpc"
	pb "github.com/dijonron/recipe-box/proto/grpc/tenant"
	tenant "github.com/dijonron/recipe-box/services/tenant/internal"

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

type tenantServer struct {
	tenant tenant.TenantManager
	pb.UnimplementedTenantServer
}

var _ pb.TenantServer = (*tenantServer)(nil)

func NewTenantServer(cfg grpc.ServerConfig, t tenant.TenantManager) Server {
	if t == nil {
		slog.Error("tenant manager is nil")
		return nil
	}

	var registers []grpc.Register
	h := &tenantServer{
		tenant: t,
	}
	registers = append(registers, h)

	s := grpc.NewServer(cfg, registers)

	as := server{s}

	return as

}

func (s *tenantServer) Register(grpcServer *g.Server) {
	pb.RegisterTenantServer(grpcServer, s)
}
