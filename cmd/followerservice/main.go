package main

import (
	"log"

	"followerservice/internal/service"
	"followerservice/pkg/gen"
	"followerservice/pkg/models"
)

func main() {
	log.Println("Starting...")

	userUpdates := make(chan models.UserUpdate)
	userStore := models.UserStore{}
	followerService := followerservice.FollowerServer{
		UnimplementedFollowerServiceServer: gen.UnimplementedFollowerServiceServer{},
		UserUpdates:                    userUpdates,
	}

	followerservice.StartUpdateManager(&userStore, userUpdates)
	followerservice.StartGrpcServer(&followerService)

	log.Println("Exiting...")
}