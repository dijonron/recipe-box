package service

import (
	"context"

	"github.com/dijonron/recipe-box/pkg/grpc"
	pb "github.com/dijonron/recipe-box/proto/grpc/user"

	g "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type userServer struct {
	pb.UnimplementedUserServer
}

var _ pb.UserServer = (*userServer)(nil)

func NewUserServer(cfg grpc.ServerConfig) Server {

	var registers []grpc.Register
	h := &userServer{}
	registers = append(registers, h)

	s := grpc.NewServer(cfg, registers)

	us := server{s}

	return us

}

func (us *userServer) Register(grpcServer *g.Server) {
	pb.RegisterUserServer(grpcServer, us)
}

func (us userServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (us userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
