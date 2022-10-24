package models

import "sync"

type User struct {
	Id         string
	Email      string
	ScreenName string
	//TODO: encapsulate these, add const
	Follows      map[string]*User
	IsFollowedBy map[string]*User
	mu           sync.Mutex
}

func (u *User) AddFollow(user *User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.Follows[user.Id] = user
}

func (u *User) AddFollowedBy(user *User) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.IsFollowedBy[user.Id] = user
}
