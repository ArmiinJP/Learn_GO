package contract

import "todolist/entity"

type UserWriteStore interface{
	Save(u entity.User) error
}

type UserLoadStore interface{
	Load() ([]entity.User, error)
}

