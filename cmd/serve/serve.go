package main

import (
	"database/sql"
	"fmt"
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
		"host=postgres port=5432 user=postgres password=postgres dbname=insta sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	postStorage := post.New(db)

	postSvc := postService.NewService(postStorage)

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	postPb.RegisterPostStorageServer(grpcServer, postSvc)

	fmt.Println("start server")
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
