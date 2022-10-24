package models 

type User struct {
	Id string
	Email string
	ScreenName string
	Follows map[string]*User
	IsFollowedBy map[string]*User
}