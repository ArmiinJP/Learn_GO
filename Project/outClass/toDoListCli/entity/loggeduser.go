package entity

import "time"

type LoggedUser struct {
	RemoteAddress string
	UserID        int
	Time          time.Time
}
