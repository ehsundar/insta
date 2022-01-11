package post

import (
	postStorage "github.com/ehsundar/insta/internal/storage/postgres/post"
	"github.com/ehsundar/insta/pkg/api/grpc/post"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func DBToProto(p *postStorage.Post) *post.Post {
	return &post.Post{
		Ident:     p.Ident,
		Caption:   p.Caption,
		Images:    p.Images,
		CreatedAt: timestamppb.New(p.CreatedAt.Time),
	}
}
