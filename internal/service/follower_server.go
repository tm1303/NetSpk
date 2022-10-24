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
	UserUpdates chan models.UserStoreAction
}

func (f *FollowerServer) CreateUser(ctx context.Context, in *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	
	newUser := models.User{
		Id: uuid.NewString(),
		Email: in.Email,
		ScreenName: in.ScreenName,
		Follows: make(map[string]*models.User),
		IsFollowedBy: make(map[string]*models.User),
	}

	response := make(chan bool)
	update := models.UserStoreAction{
		Action: func(userStore *models.UserStore){
			userStore.Append(&newUser)
			log.Printf("User store now has %v users...", userStore.Count() )
			response <- true
		},
	}

	f.UserUpdates <- update
	hasCompleted := <-response
	close(response)

	// not sure what I really to do if this fails, don't want to burn too much time :)
	if (hasCompleted){
		return &gen.CreateUserResponse{Id: newUser.Id}, nil
	}else{
		panic("failed to update!")
	}
}
