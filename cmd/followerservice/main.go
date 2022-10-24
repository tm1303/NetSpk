package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"followerservice/internal/service"
	"followerservice/pkg/gen"
)

func main() {
	log.Println("Starting...")
	setupGrpc()

	log.Println("Exiting...")
}

func setupGrpc() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := followerservice.UserServer{}

	grpcServer := grpc.NewServer()
	gen.RegisterUserServiceServer(grpcServer, &s)
	// for ease of cli dev
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
