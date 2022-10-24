package followerservice

import (
	"log"

	"followerservice/pkg/gen"
	"followerservice/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type FollowerServer struct {
	gen.UnimplementedFollowerServiceServer
	UserUpdates chan models.UserUpdate
}

func (s *FollowerServer) CreateUser(ctx context.Context, in *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	
	newUserId := uuid.NewString()
	newUser := models.User{
		Id: newUserId,
		Email: in.Email,
		ScreenName: in.ScreenName,
		Follows: make(map[string]*models.User),
		IsFollowedBy: make(map[string]*models.User),
	}

	response := make(chan bool)
	update := models.UserUpdate{
		Action: func(userStore models.UserStore){
			log.Printf("4 Running action on chan for email: %s", in.Email)
			response <- true
			return;
		},
		User: newUser,
	}

	s.UserUpdates <- update
	hasCompleted := <-response
	close(response)

	// not sure what I really to do if this fails, don't want to burn too much time :)
	if (hasCompleted){
		return &gen.CreateUserResponse{Id: newUser.Id}, nil
	}else{
		panic("failed to update!")
	}
}
