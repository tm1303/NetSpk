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
	userStoreActionQueue chan models.UserStoreAction
}

func NewFollowerServer(userStoreActionQueue chan models.UserStoreAction) *FollowerServer {
	return &FollowerServer{
		UnimplementedFollowerServiceServer: gen.UnimplementedFollowerServiceServer{},
		userStoreActionQueue:               userStoreActionQueue,
	}
}

func (f *FollowerServer) CreateUser(ctx context.Context, in *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {

	// TODO: verify no duplicates
	email := in.Email
	screenName := in.ScreenName

	newUser := &models.User{
		Id:           uuid.NewString(),
		Email:        email,
		ScreenName:   screenName,
		Follows:      make(map[string]*models.User),
		IsFollowedBy: make(map[string]*models.User),
	}

	response := make(chan bool)
	update := models.UserStoreAction{
		Action: func(userStore *models.UserStore) {
			userStore.Append(newUser)
			log.Printf("User store now has %v users...", userStore.Count())
			response <- true
		},
	}

	f.userStoreActionQueue <- update
	hasCompleted := <-response
	close(response)

	// not sure what I really to do if this fails, don't want to burn too much time :)
	if hasCompleted {
		return &gen.CreateUserResponse{Id: newUser.Id}, nil
	} else {
		panic("failed to update!")
	}
}

func (f *FollowerServer) FollowUser(ctx context.Context, in *gen.FollowUserRequest) (*gen.FollowUserResponse, error) {
	 
	user, userFound := findUser(&in.Id, f)
	followUser, followUserFound := findUser(&in.FollowId, f)

	if(!userFound){
		panic("couldnt find your user")
	}
	if(!followUserFound){
		panic("couldnt find your follow user")
	}

	response := make(chan bool)
	update := models.UserStoreAction{
		// although the User receiver funcs use a mutex I don't want anyone monkeying around with the UserStore, 
		// they might delete one of my two users out from underneith me
		Action: func(userStore *models.UserStore) {
			
			//check for nils just incase our users have since been deleted
			if(user==nil || followUser==nil){
				response <- false
			}

			user.AddFollow(followUser)
			followUser.AddFollowedBy(user)

			log.Printf("User %v now follows %v users...", user.ScreenName, followUser.ScreenName)
			response <- true
		},
	}

	f.userStoreActionQueue <- update
	hasCompleted := <-response
	close(response)

	// TODO: ALL THE ERROR HANDLING :/ (not sure what I really to do if this fails, don't want to burn too much time)
	if hasCompleted {
		return &gen.FollowUserResponse{}, nil
	} else {
		panic("failed to follow!")
	}

}

func findUser(id *string , f *FollowerServer) (*models.User, bool) {
	var user *models.User

	response := make(chan bool)
	search := models.UserStoreAction{
		// although we're not appending or deleting i don't want the store reallocated while I'm working
		Action: func(userStore *models.UserStore) {
			foundUser, found := userStore.Find(id)
			if found {
				log.Printf("found user %v", *id)
				user = foundUser
			} else {
				log.Printf("could not find user %v", *id)
			}
			response <- found
		},
	}

	f.userStoreActionQueue <- search
	userFound := <-response
	close(response)
	return user, userFound
}
