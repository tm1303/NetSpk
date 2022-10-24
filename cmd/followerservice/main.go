package main

import (
	"log"

	"followerservice/internal/service"
	"followerservice/pkg/models"
)

func main() {
	log.Println("Starting...")

	userStoreActionQueue := make(chan models.UserStoreAction)
	followerservice.StartStoreManager(models.NewUserStore(), userStoreActionQueue)

	followerService := followerservice.NewFollowerServer(userStoreActionQueue)
	followerservice.StartGrpcServer(followerService)

	log.Println("Exiting...")
}
