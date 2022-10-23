package main

import (
	"log"

	"golang.org/x/net/context"
)

type UserServer struct {
	UnimplementedUserServiceServer
}

func (s *UserServer) Create(ctx context.Context, in *CreateUserRequest) (*CreateUserResponse, error) {
	log.Printf("Receive message body from client: %s", in.Email)
	return &CreateUserResponse{Id: "new id!"}, nil
}
