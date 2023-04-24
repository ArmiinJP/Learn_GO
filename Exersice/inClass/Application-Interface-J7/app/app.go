package app

import (
	//"fmt"

	"applicationInterface/user"
)
type App struct{
	Name string
	UserStorage UserStore
}

type UserStore interface{
	CreateUser(user.User)
	ListUser() []user.User
	GetUserByID(uint) user.User
}


// type UserStore2 interface{
// 	CreateUser(user.User)
// }

// func (a App) CreateUser(u user.User) {
// 	a.UserStorage.CreateUser(u)
// }