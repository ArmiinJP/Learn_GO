package contract

import "todolist/entity"

type CategoryWriteStore interface {
	Save(c entity.Category) error
}

type CategoryLoadStore interface {
	Load(u entity.User) ([]entity.Category, error)
}
