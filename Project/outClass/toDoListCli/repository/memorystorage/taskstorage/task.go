package taskstorage

import (
	"fmt"
	"strconv"
	
	"todolist/entity"
	"todolist/constant"
)

type TaskStorage struct {
	tasks map[int][]entity.Task
}

func New() TaskStorage {
	return TaskStorage{
		tasks: make(map[int][]entity.Task),
	}
}

func (t *TaskStorage) Create(newTask entity.Task) error {
	
	t.tasks[newTask.UserID] = append(t.tasks[newTask.UserID], newTask)
	return nil
}

func (t TaskStorage) List(userID int) ([]entity.Task, error) {
	for user, userTask := range t.tasks {
		if user == userID {
			return userTask, nil
		}
	}

	return []entity.Task{}, fmt.Errorf("user doesn't have any task")
}

func (t *TaskStorage) Edit(editedTask entity.Task) error {
	for user, userTask := range t.tasks {
		if user == editedTask.UserID {
			for i, task := range userTask{
				if task.TaskID == editedTask.TaskID {
					if editedTask.CategoryID != 0 {
						task.CategoryID = editedTask.CategoryID
					}
					if editedTask.DueDate != "" {
						task.DueDate = editedTask.DueDate
					}
					if editedTask.Title != "" {
						task.Title = editedTask.Title
					}

					userTask[i] = task
				}
			}
		}
	}

	return nil
}

func (t *TaskStorage) ChangeStatus(editedTask entity.Task) error {
	for user, userTask := range t.tasks {
		if user == editedTask.UserID {
			for i, task := range userTask{
				if task.TaskID == editedTask.TaskID {
					if editedTask.IsComplete == true || editedTask.IsComplete == false {
						task.IsComplete = editedTask.IsComplete
						userTask[i] = task
					} else {
						return fmt.Errorf("value of status task is wrong this should be: true or false")
					}
				}
			}
		}
	}

	return nil
}

func (t TaskStorage) Print() {
	fmt.Println("All Task is: -------------------------")
	for user, userTask := range t.tasks {
		for _, task := range userTask{
			fmt.Printf("User ID: %d\nTask ID: %d\nCategory ID: %d\nTask Title: %s\nTask Complete: %s\nTask DueDate: %s\n",
				user, task.TaskID, task.CategoryID, task.Title, strconv.FormatBool(task.IsComplete), task.DueDate)
		}			
	}
}

func (t TaskStorage) DoesUserhaveTask(userID, taskID int) error{
	for user, userTask := range t.tasks{
		if user == userID {
			for _, task := range userTask{
				if task.TaskID == taskID{
					return nil
				}
			}
		}
	}
		
	return fmt.Errorf("user: %d doesn't have task ID: %d", userID, taskID)
}

func (t TaskStorage) NewTaskIDGenerateForUser(userID int) (int, error){
	for user, userTask := range t.tasks{
		if user == userID{
			newID := constant.MinTaskIDEachUser + len(userTask) + 1
			if newID < constant.MaxTaskIDEachUser{
				return newID, nil
			} else{
				return 0, fmt.Errorf("user dosen't allow add new task, task capacity is full")
			}
		}
	}
	return 0, nil
}
