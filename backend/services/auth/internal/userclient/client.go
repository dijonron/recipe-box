package userclient

import (
	"context"
	"errors"
	"log/slog"

	"github.com/dijonron/recipe-box/pkg/grpc"
	pb "github.com/dijonron/recipe-box/proto/grpc/user"
	"github.com/dijonron/recipe-box/services/auth/cmd/config"
	m "github.com/dijonron/recipe-box/services/auth/internal"
)

const (
	FAILED_TO_FETCH_USER string = "failed to fetch user"
)

var (
	errFailedToFetchUser = errors.New(FAILED_TO_FETCH_USER)
)

type userClient struct {
	client pb.UserClient
}

var _ m.UserClient = (*userClient)(nil)

func NewUserClient(cfg config.UserClientConfig) m.UserClient {
	conn := grpc.NewClient(cfg.Host, cfg.Port)

	client := pb.NewUserClient(conn)

	return &userClient{
		client: client,
	}
}

func (c *userClient) GetUserByEmail(ctx context.Context, email string) (m.AuthDetials, error) {

	req := &pb.GetUserByEmailRequest{
		Email: email,
	}

	resp, err := c.client.GetUserByEmail(ctx, req)
	if err != nil {
		slog.Error(FAILED_TO_FETCH_USER, "error", err)
		return m.AuthDetials{}, errFailedToFetchUser
	}

	return toAuthDetials(resp.GetUser()), nil
}

func (c *userClient) UpdateUserLogin(ctx context.Context, email string) error {
	req := &pb.UpdateUserLastLoginRequest{
		Email: email,
	}

	_, err := c.client.UpdateUserLastLogin(ctx, req)
	if err != nil {
		slog.Error(FAILED_TO_FETCH_USER, "error", err)
		return errFailedToFetchUser
	}

	return nil
}

func toAuthDetials(user *pb.UserDetails) m.AuthDetials {
	return m.AuthDetials{
		Email:        user.GetEmail(),
		PasswordHash: user.GetPasswordHash(),
	}
}
