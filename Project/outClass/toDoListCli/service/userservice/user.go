package userservice

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"todolist/delivery/requestParam"
	"todolist/delivery/responseParam"
	"todolist/entity"
)

type userRepo interface {
	AddUser(entity.User) error
	CheckDuplicateInfo(entity.User) error
	NewUserIDGenerate() (int, error) 	

	// for testing server
	//Print()
}

type Service struct {
	repository userRepo
}

func New(up userRepo) Service {
	return Service{repository: up}
}

func (s Service) RegisterUser(req requestParam.ValuesRegisterUser) (responseParam.Response, error) {
	
	hashPass := s.hashPassword([]byte(req.Password))
	
	cErr := s.repository.CheckDuplicateInfo(entity.User{Email: req.Email, Password: hashPass})
	if cErr != nil{
		return responseParam.Response{StatusCode: 400, Message: "illegal user information"}, fmt.Errorf("%s", cErr.Error())
	}
	
	NewGenerateID, nErr := s.repository.NewUserIDGenerate()
	if nErr != nil{
		return responseParam.Response{StatusCode: 500, Message: "User repository full"}, fmt.Errorf("%s", nErr.Error())
	}

	aErr := s.repository.AddUser(entity.User{
		Email: req.Email,
		Password: hashPass,
		UserID: NewGenerateID,
	})
	if aErr != nil{
		return responseParam.Response{StatusCode: 500, Message: "Faild Adding New User"}, fmt.Errorf("%s", aErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Register User Successfully", Data: []byte{}}, nil
}

func (s Service) LoginUser()  {}
func (s Service) LogoutUser() {}
func (s Service) WhichUser()  {}


func (s Service) hashPassword(password []byte) string {
	hash := sha512.New()
	hash.Write(password)

	encodedHash := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return encodedHash
}

// func (s Service) Print(){
// 	s.repository.Print()
// }