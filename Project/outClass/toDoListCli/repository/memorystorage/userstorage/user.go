package userstorage

import (
	"fmt"
	"todolist/entity"
	"todolist/constant"
)

type storage struct {
	users []entity.User
}

func New() storage {
	return storage{users: []entity.User{}}
}

func (s *storage) AddUser(u entity.User) error{
	s.users = append(s.users, u)
	return nil
}

func (s storage) CheckDuplicateInfo(u entity.User) error{
	for _, user := range s.users{
		if user.Email == u.Email && user.Password == u.Password{
			return fmt.Errorf("information is illegal, please enter new username and password")
		} 
	}

	return nil
}

func (s storage) NewUserIDGenerate() (int, error) {
	newID := constant.MinUserID + len(s.users) + 1
	if newID > constant.MaxUserID{
		return 0, fmt.Errorf("dosen't allow add new user, application capacity is full")
	}
	return newID, nil
}

func (s storage) ReturnUserID(user entity.User) (int, error){
	for _, u := range s.users{
		if u.Email == user.Email && u.Password == user.Password{
			return u.UserID, nil
		}
	}
	return 0, fmt.Errorf("user with this information not exist")
}

func (s storage) Print(){
	for _, user := range s.users{
		fmt.Printf("User ID is: %d\nEmail is: %s\nPassword is: %s \n", user.UserID, user.Email, user.Password)
	}
}