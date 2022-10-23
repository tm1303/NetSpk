package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("Starting...")

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := UserServer{}

	grpcServer := grpc.NewServer()
	RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Println("Exiting...")
}
