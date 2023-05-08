package contract

import "todolist/entity"

type TaskCreate interface {
	Save(t entity.Task) error
}

type TaskList interface {
	Load(u entity.User, c entity.Category) ([]entity.Task, error)
}


type TaskEdit interface{
	
}