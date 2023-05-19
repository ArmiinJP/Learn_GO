package requestParam

import "time"


type ValuesAddUserLoggedIn struct {
	RemoteAddress string
	UserID        int
}

type ValuesCheckedLoggedInUser struct {
	RemoteAddr string
	Time time.Time
}

type ValuesReturnLoggedInUserID struct {
	RemoteAddr string
}

type ValuesLoggedOutUser struct {
	RemoteAddr string
}
