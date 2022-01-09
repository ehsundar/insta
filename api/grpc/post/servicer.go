package post

import (
	"context"
	postStorage "github.com/ehsundar/insta/internal/storage/postgres/post"
	"github.com/ehsundar/insta/pkg/api/grpc/post"
)

type service struct {
	post.UnimplementedPostStorageServer

	storage *postStorage.Queries
}

func NewService(storage *postStorage.Queries) post.PostStorageServer {
	return &service{
		storage: storage,
	}
}

func (s *service) CreatePost(ctx context.Context, request *post.CreatePostRequest) (*post.CreatePostResponse, error) {

	newPost, err := s.storage.CreatePost(ctx, postStorage.CreatePostParams{
		Caption: request.Post.Caption,
		Images:  request.Post.Images,
	})
	if err != nil {
		return &post.CreatePostResponse{}, err
	}

	return &post.CreatePostResponse{Post: DBToProto(&newPost)}, nil
}
