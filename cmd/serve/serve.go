package main

import (
	"database/sql"
	postService "github.com/ehsundar/insta/api/grpc/post"
	postPb "github.com/ehsundar/insta/pkg/api/grpc/post"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"

	"github.com/ehsundar/insta/internal/storage/postgres/post"
)

func main() {
	db, err := sql.Open("postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=insta sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	postStorage := post.New(db)

	postSvc := postService.NewService(postStorage)

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	postPb.RegisterPostStorageServer(grpcServer, postSvc)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
