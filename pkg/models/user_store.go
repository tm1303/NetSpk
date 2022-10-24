package models

import "log"


type UserStore struct {
	users map[string]*User
}

func NewUserStore()(*UserStore){
	return &UserStore{users: make(map[string]*User)}
}

func (s *UserStore) Append(user *User) {

	log.Printf("Append email %v ...", user.Email)
	log.Printf("Append screenName %v ...", user.ScreenName)
	s.users[user.Id] = user
}

func (s *UserStore) Count() int {
	return len(s.users)
}

func (s *UserStore) Find(userId *string)(*User, bool){
	user, found := s.users[*userId]
	return user, found
}
