package memorystorage

import (
	"fmt"
	"strconv"
	"todolist/entity"
)

type TaskStorage struct {
	tasks []entity.Task
}

func (t *TaskStorage) Create(newTask entity.Task) error {
	t.tasks = append(t.tasks, newTask)

	return nil
}

func (t TaskStorage) List(userID int) ([]entity.Task, error) {
	var userTasks = []entity.Task{}
	for _, v := range t.tasks {
		if v.UserID == userID {
			userTasks = append(userTasks, v)
		}
	}

	return userTasks, nil
}

func (t *TaskStorage) Edit(editedTask entity.Task) error {
	for _, v := range t.tasks {
		if v.UserID == editedTask.UserID && v.ID == editedTask.ID {
			if editedTask.CategoryID != 0 {
				v.CategoryID = editedTask.CategoryID
			}
			if editedTask.DueDate != "" {
				v.DueDate = editedTask.DueDate
			}
			if editedTask.Title != "" {
				v.Title = editedTask.Title
			}
		}
	}

	return nil
}

func (t *TaskStorage) ChangeStatus(editedTask entity.Task) error {
	for _, v := range t.tasks {
		if v.UserID == editedTask.UserID && v.ID == editedTask.ID {
			v.IsComplete = editedTask.IsComplete
		}
	}

	return nil
}

func (t TaskStorage) Print() {
	fmt.Println("All Task is: -------------------------")
	for _, v := range t.tasks {
		fmt.Printf("User ID: %d\nTask ID: %d\nCategory ID: %d\nTask Title: %s\nTask Complete: %s\nTask DueDate: %s\n",
			v.UserID, v.ID, v.CategoryID, v.Title, strconv.FormatBool(v.IsComplete), v.DueDate)
	}
}
