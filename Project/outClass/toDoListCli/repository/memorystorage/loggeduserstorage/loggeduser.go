package loggeduserstorage

import (
	"fmt"
	"time"
	"todolist/entity"
)

type storage struct{
	loggedUsers map[string]entity.LoggedUser
}

func New() storage {
	return storage{loggedUsers: make(map[string]entity.LoggedUser)}
}

func (s *storage) AddUser(user entity.LoggedUser) error{
	s.loggedUsers[user.RemoteAddress] = user
	return nil
}

func (s storage) CheckUser(user entity.LoggedUser) error{
	userLoggedInfo, ok := s.loggedUsers[user.RemoteAddress]
	if !ok {
		return fmt.Errorf("user not login")
	}
	
	if user.Time.Before(userLoggedInfo.Time){
		return nil
	} else {
		return fmt.Errorf("login time is Expired, please login again")
	}
}

func (s storage) ReturnUserID (user entity.LoggedUser) (int, error){
	
	return  s.loggedUsers[user.RemoteAddress].UserID, nil
}

func (s storage) CheckTimeExpired(t time.Time) []string{

	var loggedUserTimeExpired = []string{}

	for userAddr, userInfo := range s.loggedUsers{
		if t.Before(userInfo.Time){
			loggedUserTimeExpired = append(loggedUserTimeExpired, userAddr)
		}
	}

	return loggedUserTimeExpired
}

func (s *storage) DeleteUser(userAddr string) error{
	
	delete(s.loggedUsers, userAddr)
	return nil
}

func (s storage) Print(){
	for _, userinfo := range s.loggedUsers {
		fmt.Printf("User ID is: %d\nAddress is: %s\nTime Expired is: %s \n", userinfo.UserID, userinfo.RemoteAddress, userinfo.Time.String())
	}
}