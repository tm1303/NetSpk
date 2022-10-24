package followerservice

import (
	"log"

	"golang.org/x/net/context"
	"followerservice/pkg/gen"
)

type UserServer struct {
	gen.UnimplementedUserServiceServer
}

func (s *UserServer) Create(ctx context.Context, in *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	log.Printf("Receive message body from client: %s", in.Email)
	return &gen.CreateUserResponse{Id: "new id!"}, nil
}
