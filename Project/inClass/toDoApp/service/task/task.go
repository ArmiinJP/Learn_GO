package task

import (
	"fmt"
	"todolistapp/entity"
)

type ServiceRepository interface {
	//DoesUserhaveCategory(userID, categoryID int) bool
	CreateNewTask(entity.Task) (entity.Task, error)
	ListUserTask(userID int) ([]entity.Task, error)
}

type Service struct {
	repository ServiceRepository
}

func NewService(repo ServiceRepository) Service{
	return Service{repository: repo}
}

type CreatedRequest struct {
	Title              string
	DueDate            string
	CategoryID         int
	AutheticatedUserID int
}

type CreatedResponse struct {
	Task entity.Task
}

func (t Service) CreatedTask(req CreatedRequest) (CreatedResponse, error) {

	// if !t.repository.DoesUserhaveCategory(req.AutheticatedUserID, req.CategoryID) {
	// 	return CreatedResponse{}, fmt.Errorf("user have this category %d", req.CategoryID)
	// }

	createdTask, cErr := t.repository.CreateNewTask(entity.Task{
		Title:      req.Title,
		DueDate:    req.DueDate,
		CategoryID: req.CategoryID,
		IsDone:     false,
		UserID:     req.AutheticatedUserID})

	if cErr != nil {
		return CreatedResponse{}, fmt.Errorf("can't create new Task %v", cErr)
	}

	return CreatedResponse{Task: createdTask}, nil
}

type ListRequest struct {
	UserID int
}

type listResponse struct {
	Tasks []entity.Task
}

func (t Service) ListTask(req ListRequest) (listResponse, error) {
	tasks, err := t.repository.ListUserTask(req.UserID)
	if err != nil {
		return listResponse{}, fmt.Errorf("error! %v", err)
	}

	return listResponse{Tasks: tasks}, nil
}
