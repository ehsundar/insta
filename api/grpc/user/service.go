package user

import (
	"context"
	userStorage "github.com/ehsundar/insta/internal/storage/postgres/user"
	"github.com/ehsundar/insta/pkg/api/grpc/user"
)

type service struct {
	user.UnimplementedUserStorageServer

	storage *userStorage.Queries
}

func NewService(storage *userStorage.Queries) user.UserStorageServer {
	return &service{
		storage: storage,
	}
}

func (s *service) CreateUser(ctx context.Context, request *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	ident, err := s.storage.CreateUser(ctx, userStorage.CreateUserParams{
		Username: request.Username,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		Ident: ident,
	}, nil
}

func (s *service) GetUserByIdent(ctx context.Context, request *user.GetUserByIdentRequest) (*user.GetUserByIdentResponse, error) {
	u, err := s.storage.GetUserByIdent(ctx, request.Ident)
	if err != nil {
		return nil, err
	}

	return &user.GetUserByIdentResponse{User: DBToProto(&u)}, nil
}
