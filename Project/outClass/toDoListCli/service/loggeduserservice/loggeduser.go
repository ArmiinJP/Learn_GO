package loggeduserservice

import (
	"fmt"
	"time"
	"todolist/delivery/requestParam"
	"todolist/delivery/responseParam"
	"todolist/entity"
)

type loggedRepo interface {
	AddUser(entity.LoggedUser) error
	CheckUser(entity.LoggedUser) error
	ReturnUserID (entity.LoggedUser) (int, error)
	CheckTimeExpired(time.Time) []string
	DeleteUser(string) error

	// for testing server
	Print()		
}

type Service struct {
	repository loggedRepo
}

func New(lr loggedRepo) Service{
	return Service{repository: lr}
}

func (s Service) AddLoggedInUser(request requestParam.ValuesAddUserLoggedIn) (responseParam.Response, error) {
	aErr := s.repository.AddUser(entity.LoggedUser{
		RemoteAddress: request.RemoteAddress,
		UserID:        request.UserID,
		Time:    		time.Now().Local().Add(time.Minute * time.Duration(20)),
	})
	if aErr != nil {
		return responseParam.Response{StatusCode: 500, Message: "Failed logged in User"}, fmt.Errorf("error login user in loggedUserService %s", aErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Login was successful, Expiration time 20 minutes", Data: []byte{}}, nil
}

func (s Service) CheckLoggedINUser(request requestParam.ValuesCheckedLoggedInUser) (responseParam.Response, error) {
	cErr := s.repository.CheckUser(entity.LoggedUser{
		RemoteAddress: request.RemoteAddr,
		Time: request.Time,
	})
	if cErr != nil{
		
		return responseParam.Response{StatusCode: 400, Message: "Please login first!", Data: []byte{}}, fmt.Errorf("user is not Logged in")
	}
	return responseParam.Response{}, nil
}

func (s Service) ReturnLoggedInUserID(request requestParam.ValuesReturnLoggedInUserID) (int, responseParam.Response, error){
	userID, rErr := s.repository.ReturnUserID(entity.LoggedUser{RemoteAddress: request.RemoteAddr,})
	if rErr != nil{
		return 0, responseParam.Response{StatusCode: 500, Message: "Failed to logging User", Data: []byte{}} ,fmt.Errorf("error return userID in logged in process")
	}
	return userID, responseParam.Response{}, nil
}

func (s Service) CheckLoggedInTimeExpired(t time.Time) []string{
	userTimeExpired := s.repository.CheckTimeExpired(t)
	return userTimeExpired
}

func (s Service) LoggedOutUser(userRemoteAddr string) error{
	dErr := s.repository.DeleteUser(userRemoteAddr)
	if dErr != nil{
		return fmt.Errorf("error logged out user with address: %s", userRemoteAddr)
	}

	return nil
}

func (s Service) Print(){
	s.repository.Print()
}