package taskservice

import (
	"encoding/json"
	"fmt"
	"todolist/delivery/requestParam"
	"todolist/delivery/response"
	"todolist/entity"
)

type TaskRepo interface {
	Create(task entity.Task) error
	List(userID int) ([]entity.Task, error)
	Edit(task entity.Task) error 
	ChangeStatus(editedTask entity.Task) error
}

type Service struct {
	repository TaskRepo
}

func New(tr TaskRepo) Service{
	return Service{repository: tr}
}

func (s Service) CreateTaskRequest(request requestParam.ValuesCreateTask) (responseParam.Response, error){
	cErr := s.repository.Create(entity.Task{
		Title: request.Title,
		DueDate: request.DueDate,
		CategoryID: request.CategoryID,
		IsComplete: false,
		UserID: request.UserID,
	})

	if cErr != nil{
		return responseParam.Response{StatusCode: 500, Message: "Failed to Create Task", Data: []byte{}}, fmt.Errorf("error Creating task: %s", cErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Create Task Successfully", Data:[]byte{}}, nil
}

func (s Service) ListTaskRequest(request requestParam.ValuesListTask) (responseParam.Response, error){
	
	userTask, lErr := s.repository.List(request.UserID)
	if lErr != nil{
	
		return responseParam.Response{StatusCode: 500, Message: "Failed to List Task", Data: []byte{}}, fmt.Errorf("error Listing task: %s", lErr.Error())
	}

	//fmt.Println(userTask)
	data, mErr := json.Marshal(userTask)
	if mErr != nil{
		
		return responseParam.Response{StatusCode: 500, Message: "Failed to List Task", Data: []byte{}}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}
	
	return responseParam.Response{StatusCode: 200, Message: "Return List Task Successfully", Data: data}, nil
}	

func (s Service) EditTaskRequst(request requestParam.ValuesEditTask) (responseParam.Response, error){
	eErr := s.repository.Edit(entity.Task{
		ID: request.ID,
		Title: request.Title,
		DueDate: request.DueDate,
		CategoryID: request.CategoryID,
		IsComplete: request.IsComplete,
		UserID: request.UserID,
	})

	if eErr != nil{
		return responseParam.Response{StatusCode: 500, Message: "Failed to Edit Task", Data: []byte{}}, fmt.Errorf("error Editing task: %s", eErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Edit Task Successfully", Data:[]byte{}}, nil
}

func (s Service) ListTodayRequest(request requestParam.ValueslistTodayTask) (responseParam.Response, error){
	return responseParam.Response{}, nil
}

func (s Service) ListDayRequest(request requestParam.ValuesListSpecificDayTask)(responseParam.Response, error){
	return responseParam.Response{}, nil
}

func (s Service) ChangeStatusRequest(request requestParam.ValuesChangeStatusTask)(responseParam.Response, error){
	return responseParam.Response{}, nil
}


