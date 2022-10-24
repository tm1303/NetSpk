package followerservice

import (
	"followerservice/pkg/models"
)

// maintain single access to the store
func StartStoreManager(userStore *models.UserStore, userUpdates chan models.UserStoreAction) {
	go func() {
		for {
			nextUpdate := <-userUpdates
			nextUpdate.Action(userStore)
		}
	}()
}