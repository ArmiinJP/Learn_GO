package memorystore

import (
	"todolistapp/entity"
)

type Task struct {
	tasks []entity.Task
}

func NewTask() *Task{
 	return &Task{tasks: make([]entity.Task, 0)}
}

func (t Task) ListUserTask(userID int) ([]entity.Task, error) {

	var userTasks []entity.Task
	for _, task := range t.tasks {
		if task.UserID == userID {
			userTasks = append(userTasks, task)
		}
	}
	return userTasks, nil

}

func (t *Task) CreateNewTask(task entity.Task) (entity.Task, error) {
	task.ID = len(t.tasks) + 1
	t.tasks = append(t.tasks, task)

	return task, nil
}
