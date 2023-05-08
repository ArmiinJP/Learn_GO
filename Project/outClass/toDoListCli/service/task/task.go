package taskservice

import (
	"fmt"
	"todolist/delivery/requestParam"
	"todolist/delivery/response"
	"todolist/entity"
)

type TaskRepo interface {
	Create(entity.Task) error
	//List(u entity.User, c entity.Category) ([]entity.Task, error)
	//Edit(entity.Task, entity.Task) error 
}

type Service struct {
	repository TaskRepo
}

func New(tr TaskRepo) Service{
	return Service{repository: tr}
}

func (s Service) CreateTaskRequest(request requestParam.ValuesCreateTask) (response.Param, error){
	cErr := s.repository.Create(entity.Task{
		Title: request.Title,
		DueDate: request.DueDate,
		CategoryID: request.CategoryID,
		IsComplete: false,
		UserID: request.UserID,
	})

	if cErr != nil{
		return response.Param{StatusCode: 500, Message: "Failed to Create Task"}, fmt.Errorf("error Creating task: %s", cErr.Error())
	}

	return response.Param{StatusCode: 200, Message: "Create Task Successfully"}, nil
}

func ListTaskRequest(){}
func EditRequst(){}
func ListTodayRequest(){}
func ListDayRequest(){}
func ChangeStatusRequest(){}


