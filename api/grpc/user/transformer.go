package user

import (
	userStorage "github.com/ehsundar/insta/internal/storage/postgres/user"
	"github.com/ehsundar/insta/pkg/api/grpc/user"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DBToProto(u *userStorage.User) *user.User {
	return &user.User{
		Ident:     u.Ident,
		Username:  u.Username,
		CreatedAt: timestamppb.New(u.CreatedAt.Time),
	}
}
