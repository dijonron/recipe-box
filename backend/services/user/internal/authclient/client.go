package userclient

import (
	"context"
	"errors"
	"log/slog"

	"github.com/dijonron/recipe-box/pkg/grpc"
	pb "github.com/dijonron/recipe-box/proto/grpc/auth"
	"github.com/dijonron/recipe-box/services/user/cmd/config"
	m "github.com/dijonron/recipe-box/services/user/internal"
)

const (
	FAILED_TO_AUTHENTICATE_USER string = "failed to authenticate user"
)

var (
	errFailedToAuthenticateUser = errors.New(FAILED_TO_AUTHENTICATE_USER)
)

type authClient struct {
	client pb.AuthClient
}

var _ m.AuthClient = (*authClient)(nil)

func NewAuthClient(cfg config.AuthClientConfig) m.AuthClient {
	conn := grpc.NewClient(cfg.Host, cfg.Port)

	client := pb.NewAuthClient(conn)

	return &authClient{
		client: client,
	}
}

func (c *authClient) AuthenticateUser(ctx context.Context, email, password string) (string, error) {
	req := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	resp, err := c.client.Login(ctx, req)
	if err != nil {
		slog.Error(FAILED_TO_AUTHENTICATE_USER, "error", err)
		return "", errFailedToAuthenticateUser
	}

	return resp.GetToken(), nil
}
