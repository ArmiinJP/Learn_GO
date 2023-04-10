package user

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"ErrorAndLog/richError"
	"ErrorAndLog/simpleError"
)

type User struct{
	Id int
	Name string
}

func (u User) String() string{
	return fmt.Sprintf("{ID is: %d, Name is: %s}", u.Id, u.Name)
}

func (u User) CheckUserReturnRichErr() (User, error){
	if u.Id <= 0 {
		return User{}, richError.RichError{
			Message: "id is not valid",
			Operation: "CheckUserRich",
			MetaData: map[string]string{
				"info": "user_Data",
				"id": strconv.Itoa(u.Id),
				"name": u.Name,
			},Time: time.Now(),
		}
	} else {
		return u, nil
	}
}

func (u User) CheckUserReturnSimpleErr() (User, error){
	if u.Id <= 0 {
		return User{}, simpleError.SimpleError{
			Message: "id is not valid",
			Operation: "CheckUserSimple",
		}
	} else {
		return u, nil
	}
}

func (u User) CheckUserReturnFmtErrorf() (User, error){
	if u.Id <= 0 {
		return User{}, fmt.Errorf("user invalid")
	} else {
		return u, nil
	}
}

func (u User) CheckUserReturnErrorsNew() (User, error){
	if u.Id <= 0 {
		return User{}, errors.New("user invalid")
	} else {
		return u, nil
	}
}