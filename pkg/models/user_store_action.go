package models 


type UserStoreAction struct {
	Action func(userStore *UserStore)
}
