package memorystorage

import (
	"fmt"
	"todolist/entity"
)

type TaskStorage struct{
	tasks []entity.Task
}

func (t *TaskStorage) Create(newTask entity.Task) error{
	t.tasks = append(t.tasks, newTask)
	
	return nil
}

func (t TaskStorage) Print(){
	fmt.Println(t.tasks)
}
//func (t TaskStorage) List(user entity.User)