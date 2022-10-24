package models 


type UserUpdate struct {
	Action func(userStore *UserStore)
}
