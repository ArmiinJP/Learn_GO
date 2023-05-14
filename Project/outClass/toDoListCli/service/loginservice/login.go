package loginservice

import "time"

type Login struct{
	users map[int]string
	expireTime time.Time
}