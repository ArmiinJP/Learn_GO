package main

import (
	"fmt"

	"applicationInterface/app"
	"applicationInterface/storage"
	"applicationInterface/user"
)

func main(){
	
	var newApp = app.App{
		Name: "testApp",
		UserStorage: &storage.InMemoryMap{},
		//UserStorage: &storage.InMemoryMap{},
	}

	
	var newUser, newUser2 = user.User{
		Name: "hasan",
		ID: 1,
	},user.User{
		Name: "hossain",
		ID: 2,
	}

	newApp.UserStorage.CreateUser(newUser)
	newApp.UserStorage.CreateUser(newUser2)
	
	fmt.Println(newApp.UserStorage.ListUser())

}
