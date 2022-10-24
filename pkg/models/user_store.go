package models


type UserStore struct {
	users []*User
}

func (s *UserStore) Append(user *User) {
	s.users = append(s.users, user)
}

func (s *UserStore) Count()(int){
	return len( s.users)
}
