package followerservice

import (
	"followerservice/pkg/models"
)

func StartUpdateManager(userStore *models.UserStore, userUpdates chan models.UserStoreAction) {
	go func() {
		for {
			nextUpdate := <-userUpdates
			nextUpdate.Action(userStore)
		}
	}()
}