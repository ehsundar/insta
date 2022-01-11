package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	postService "github.com/ehsundar/insta/api/grpc/post"
	userService "github.com/ehsundar/insta/api/grpc/user"
	"github.com/ehsundar/insta/internal/storage/postgres/post"
	"github.com/ehsundar/insta/internal/storage/postgres/user"
	postPb "github.com/ehsundar/insta/pkg/api/grpc/post"
	userPb "github.com/ehsundar/insta/pkg/api/grpc/user"
)

func main() {
	db, err := sql.Open("postgres",
		"host=postgres port=5432 user=postgres password=postgres dbname=insta sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	postStorage := post.New(db)
	userStorage := user.New(db)

	postSvc := postService.NewService(postStorage)
	userSvc := userService.NewService(userStorage)

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	postPb.RegisterPostStorageServer(grpcServer, postSvc)
	userPb.RegisterUserStorageServer(grpcServer, userSvc)

	fmt.Println("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
