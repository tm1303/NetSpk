package followerservice

import (
	"followerservice/pkg/models"
)

func StartUpdateManager(userStore models.UserStore, userUpdates chan models.UserUpdate) {
	go func() {
		for {
			nextUpdate := <-userUpdates
			nextUpdate.Action(userStore)
		}
	}()
}