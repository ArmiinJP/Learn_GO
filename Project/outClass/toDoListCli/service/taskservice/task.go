package taskservice

import (
	"encoding/json"
	"fmt"
	"time"
	"todolist/delivery/requestParam"
	"todolist/delivery/responseParam"
	"todolist/entity"
)

type taskRepo interface {
	Create(task entity.Task) error
	List(userID int) ([]entity.Task, error)
	Edit(task entity.Task) error
	ChangeStatus(editedTask entity.Task) error
	DoesUserhaveTask(userID, taskID int) error
	NewTaskIDGenerateForUser(userID int) (int, error) 

	// for testing server
	//Print()
}

type Service struct {
	repository taskRepo
}

func New(tr taskRepo) Service {
	return Service{repository: tr}
}

func (s Service) CreateTaskRequest(request requestParam.ValuesCreateTask) (responseParam.Response, error) {
	
	NewTaskIDGenerate, nErr := s.repository.NewTaskIDGenerateForUser(request.UserID)
	if nErr != nil{
		return responseParam.Response{StatusCode: 400, Message: "User Task capacity is full"}, fmt.Errorf("%s", nErr.Error())
	}

	cErr := s.repository.Create(entity.Task{
		TaskID: NewTaskIDGenerate,
		Title:      request.Title,
		DueDate:    request.DueDate,
		CategoryID: request.CategoryID,
		IsComplete: false,
		UserID:     request.UserID,
	})
	if cErr != nil {
		return responseParam.Response{StatusCode: 500, Message: "Failed Creating new Task"}, fmt.Errorf("error Creating task: %s", cErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Create Task Successfully", Data: []byte{}}, nil
}

func (s Service) ListTaskRequest(request requestParam.ValuesListTask) (responseParam.Response, error) {

	userTask, lErr := s.repository.List(request.UserID)
	if lErr != nil {

		return responseParam.Response{StatusCode: 200, Message: "User doesn't have any Task" }, fmt.Errorf("error Listing task: %s", lErr.Error())
	}

	//fmt.Println(userTask)
	data, mErr := json.Marshal(userTask)
	if mErr != nil {

		return responseParam.Response{StatusCode: 500, Message: "Faild to Listing Task"}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Return List Task Successfully", Data: data}, nil
}

func (s Service) ListTodayRequest(request requestParam.ValueslistTodayTask) (responseParam.Response, error) {

	userTask, lErr := s.repository.List(request.UserID)
	if lErr != nil {

		return responseParam.Response{StatusCode: 404, Message: "User doesn't have any Task" }, fmt.Errorf("error Listing task: %s", lErr.Error())
	}

	var todayTask = []entity.Task{}
	for _, v := range(userTask){
		date, _ := time.Parse("2006-01-02", v.DueDate)
		if (date.Year() == request.Date.Year()) && (int(date.Month()) == int(request.Date.Month())) && (date.Day() == request.Date.Day()){
			todayTask = append(todayTask, v)
		}
	}

	if len(todayTask) == 0 {

		return responseParam.Response{StatusCode: 200, Message: "User doesn't have any task today"}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}

	data, mErr := json.Marshal(todayTask)
	if mErr != nil {

		return responseParam.Response{StatusCode: 500, Message: "Failed listing today task"}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Return List Today Task Successfully", Data: data}, nil
}	

func (s Service) ListSpecificDayRequest(request requestParam.ValuesListSpecificDayTask) (responseParam.Response, error) {


	userTask, lErr := s.repository.List(request.UserID)
	if lErr != nil {

		return responseParam.Response{StatusCode: 404, Message: "User doesn't have any Task" }, fmt.Errorf("error Listing task: %s", lErr.Error())
	}

	var dayTask = []entity.Task{}
	for _, v := range(userTask){
		date, _ := time.Parse("2006-01-02", v.DueDate)
		if (date.Year() == request.Date.Year()) && (int(date.Month()) == int(request.Date.Month())) && (date.Day() == request.Date.Day()){
			dayTask = append(dayTask, v)
		}
	}

	if len(dayTask) == 0 {

		return responseParam.Response{StatusCode: 200, Message: "User doesn't have any task this day"}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}

	data, mErr := json.Marshal(dayTask)
	if mErr != nil {

		return responseParam.Response{StatusCode: 500, Message: "Failed listing this task"}, fmt.Errorf("error Marshaling Response: %s", lErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Return List Specific Day Task Successfully", Data: data}, nil
}	

func (s Service) EditTaskRequst(request requestParam.ValuesEditTask) (responseParam.Response, error) {
	
	dErr := s.repository.DoesUserhaveTask(request.UserID, request.TaskID)
	if dErr != nil{
		return responseParam.Response{StatusCode: 400, Message: "Dosn't have this Task"}, fmt.Errorf("user doesn't have Task ID: %d error: %s",request.TaskID ,dErr.Error())
	}

	eErr := s.repository.Edit(entity.Task{
		TaskID:     request.TaskID,
		Title:      request.Title,
		DueDate:    request.DueDate,
		CategoryID: request.CategoryID,
		IsComplete: request.IsComplete,
		UserID:     request.UserID,
	})

	if eErr != nil {
		return responseParam.Response{StatusCode: 500, Message: "Failed to Edit Task"}, fmt.Errorf("error Editing task: %s", eErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Edit Task Successfully", Data: []byte{}}, nil
}

func (s Service) ChangeStatusRequest(request requestParam.ValuesChangeStatusTask) (responseParam.Response, error) {
	
	dErr := s.repository.DoesUserhaveTask(request.UserID, request.TaskID)
	if dErr != nil{
		return responseParam.Response{StatusCode: 400, Message: "Dosn't have this Task"}, fmt.Errorf("user doesn't have Task ID: %d error: %s",request.TaskID ,dErr.Error())
	}
	
	cErr := s.repository.ChangeStatus(entity.Task{
		TaskID: request.TaskID,
		IsComplete: request.IsComplete,
		UserID: request.UserID,
	})
	if cErr != nil{
		return responseParam.Response{StatusCode: 500, Message: "Failed to change status Task"}, fmt.Errorf("error  Change status Task: %s", cErr.Error())
	}

	return responseParam.Response{StatusCode: 200, Message: "Change Status Task Successfully", Data: []byte{}}, nil
}

// func (s Service) Print(){
// 	s.repository.Print()
// }