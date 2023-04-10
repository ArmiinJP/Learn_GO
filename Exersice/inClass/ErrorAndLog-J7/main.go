package main

import (
	"fmt"
	"os"

	"ErrorAndLog/log"
	"ErrorAndLog/user"
)

//import "errors"



func main () {

	var l = log.Log{}

	var u1 = user.User{
		Id: -123,
		Name: "ali",
	}

	if u, rErr := u1.CheckUserReturnRichErr(); rErr != nil{
		fmt.Println(rErr.Error())
		l.Add(rErr)
	} else {
		fmt.Println(u)
	}

	if u, sErr := u1.CheckUserReturnSimpleErr(); sErr != nil{
		fmt.Println(sErr.Error())
		l.Add(sErr)
	} else {
		fmt.Println(u)
	}

	if u, sErr := u1.CheckUserReturnFmtErrorf(); sErr != nil{
		fmt.Println(sErr.Error())
		l.Add(sErr)
	} else {
		fmt.Println(u)
	}
	
	if u, sErr := u1.CheckUserReturnErrorsNew(); sErr != nil{
		fmt.Println(sErr.Error())
		l.Add(sErr)
	} else {
		fmt.Println(u)
	}	

	if _, tErr := os.OpenFile("alalki", os.O_RDONLY, 0666); tErr != nil{
		fmt.Println("error occured: ", tErr.Error())
		l.Add(tErr)
	}

	l.Save()
}

