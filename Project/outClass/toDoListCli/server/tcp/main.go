package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	//"os"
	"strings"

	"todolist/contract"
	"todolist/delivery/requestParam"
	"todolist/delivery/responseParam"
	"todolist/entity"
	"todolist/repository/filestorage"
	"todolist/repository/memorystorage/categorystorage"
	"todolist/repository/memorystorage/taskstorage"
	"todolist/repository/memorystorage/userstorage"
	"todolist/service/categoryservice"
	"todolist/service/taskservice"
	"todolist/service/userservice"
)

var (
	users []entity.User
)

func main() {
	userAddressflag := flag.String("address", ":2023", "enter server Address(IP:PORT)")
	serialFlagUser := flag.String("serialized", "", "enter your format to save file")
	flag.Parse()

	serializationMode := parsingFlag(*serialFlagUser)

	//just change this assignemnt
	var storageFile = filestorage.New(serializationMode)

	//var writeUser contract.UserWriteStore = storageFile
	var loadUser contract.UserLoadStore = storageFile

	if usersStorage, err := loadUser.Load(); err == nil {
		users = append(users, usersStorage...)
	}

	var userRepo = userstorage.New()
	var categoryRepo = categorystorage.New()
	var taskRepo = taskstorage.New()
	
	
	var userService = userservice.New(&userRepo)
	var categoryService = categoryservice.New(&categoryRepo)
	var taskService = taskservice.New(&taskRepo)
	

	NetworkLayer(*userAddressflag, taskService, categoryService, userService)
}

func parsingFlag(serialFlag string) string {

	// parsing serialFlag
	switch strings.ToLower(serialFlag) {
	case "json", "xml", "csv", "txt":
		serialFlag = strings.ToLower(serialFlag)

	default:
		fmt.Println("Format File Not determine or False")
		serialFlag = "json"
	}

	return serialFlag
}

func NetworkLayer(ListeningAddressPort string, tService taskservice.Service, cService categoryservice.Service, uService userservice.Service) {
	listener, lErr := net.Listen("tcp", ListeningAddressPort)
	if lErr != nil {
		log.Fatalln("listening to the port refused: ", lErr.Error())
	}
	defer listener.Close()

	for {
		conn, aErr := listener.Accept()
		if aErr != nil {
			log.Println("Accept Connection Error: ", aErr.Error())

			continue
		}

		var req = make([]byte, 1024)
		numberOfByte, rErr := conn.Read(req)
		if rErr != nil {
			log.Println("Reading Data Error: ", rErr.Error())
			
			continue
		}

		var request = requestParam.Request{}
		uErr := json.Unmarshal(req[:numberOfByte], &request)
		if uErr != nil {
			log.Printf("Unmarshaling request error : %s", uErr.Error())

			continue
		}

		response, pErr := processRequest(request, tService, cService, uService)
		if pErr != nil {

			log.Println("Error Processing request: ", pErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			log.Printf("error Marshaling response %s", mErr.Error())
		}

		_, wErr := conn.Write(dataResponse)
		if wErr != nil {
			log.Println("Error responsing Client: ", wErr.Error())

			continue
		}
		
		//uService.Print()
		//tService.Print()
		//cService.Print()
	}
}

func processRequest(req requestParam.Request, tService taskservice.Service, cService categoryservice.Service, uService userservice.Service) (responseParam.Response, error) {
	switch req.Command {
	case "create-task":
		createTaskRequestParam := &requestParam.ValuesCreateTask{}
		uErr := json.Unmarshal(req.ValueCommand, createTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter Create Task"}, fmt.Errorf("error unmarshaling value-create-request: %s", uErr.Error())
		}

		response, cErr := cService.CheckCategoryID(createTaskRequestParam.UserID, createTaskRequestParam.CategoryID)
		if cErr != nil {
			return response, fmt.Errorf(cErr.Error())
		}

		response, cErr = tService.CreateTaskRequest(*createTaskRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service create-task-request: %s", cErr.Error())
		}

		return  response, nil
	case "list-task":
		listTaskRequestParam := &requestParam.ValuesListTask{}
		uErr := json.Unmarshal(req.ValueCommand, listTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter List Task"}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		response, cErr := tService.ListTaskRequest(*listTaskRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service list-task-request: %s", cErr.Error())
		}

		return response, nil
	case "list-today-task":
		listTodayTaskRequestParam := &requestParam.ValueslistTodayTask{}
		uErr := json.Unmarshal(req.ValueCommand, listTodayTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter List today Task"}, fmt.Errorf("error unmarshaling value-list-today-request: %s", uErr.Error())
		}

		response, cErr := tService.ListTodayRequest(*listTodayTaskRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service list-today-task-request: %s", cErr.Error())
		}

		return response, nil
	case "list-day-task":
		listSpecificDayTaskRequestParam := &requestParam.ValuesListSpecificDayTask{}
		uErr := json.Unmarshal(req.ValueCommand, listSpecificDayTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter List specific Day Task"}, fmt.Errorf("error unmarshaling value-list-specific-day-request: %s", uErr.Error())
		}

		response, cErr := tService.ListSpecificDayRequest(*listSpecificDayTaskRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service list-specific-day-task-request: %s", cErr.Error())
		}

		return response, nil
	case "edit-task":
		editTaskRequestParam := &requestParam.ValuesEditTask{}
		uErr := json.Unmarshal(req.ValueCommand, editTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter edit Task"}, fmt.Errorf("error unmarshaling value-edit-request: %s", uErr.Error())
		}

		response, cErr := cService.CheckCategoryID(editTaskRequestParam.UserID, editTaskRequestParam.CategoryID)
		if cErr != nil {
			return response, fmt.Errorf(cErr.Error())
		}

		response, cErr = tService.EditTaskRequst(*editTaskRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service edit-task-request: %s", cErr.Error())
		}

		return response, nil
	case "change-status-task":
		changeStatusTaskReuquestParam := &requestParam.ValuesChangeStatusTask{}
		uErr := json.Unmarshal(req.ValueCommand, changeStatusTaskReuquestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter change status Task"}, fmt.Errorf("error unmarshaling value-change-status-request: %s", uErr.Error())
		}

		response, cErr := tService.ChangeStatusRequest(*changeStatusTaskReuquestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service change-status-task-request: %s", cErr.Error())
		}

		return response, nil
	case "create-category":
		createCategoryRequestParam := &requestParam.ValuesCreateCategory{}
		uErr := json.Unmarshal(req.ValueCommand, createCategoryRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter create Category"}, fmt.Errorf("error unmarshaling value-create-category-request: %s", uErr.Error())
		}

		response, cErr := cService.CreateCategoryRequest(*createCategoryRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service create-category-request: %s", cErr.Error())
		}

		return response, nil
	case "list-category":
		listCategoryRequestParam := &requestParam.ValuesListCategory{}
		uErr := json.Unmarshal(req.ValueCommand, listCategoryRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter List Category"}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		response, cErr := cService.ListCategoryRequest(requestParam.ValuesListCategory(*listCategoryRequestParam))
		if cErr != nil {

			return response, fmt.Errorf("error in service list-task-request: %s", cErr.Error())
		}

		return response, nil
	case "edit-category":
		editCategoryRequestParam := &requestParam.ValuesEditCategory{}
		uErr := json.Unmarshal(req.ValueCommand, editCategoryRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter edit Category"}, fmt.Errorf("error unmarshaling value-edit-category-request: %s", uErr.Error())
		}

		response, cErr := cService.EditCategoryRequst(*editCategoryRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service edit-category-request: %s", cErr.Error())
		}

		return response, nil
	case "register-user":
		registerUserRequest := &requestParam.ValuesRegisterUser{}
		uErr := json.Unmarshal(req.ValueCommand, registerUserRequest)
		if uErr != nil{

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter Register User"}, fmt.Errorf("error unmarshaling register-user-request: %s", uErr.Error())
		}

		response, rErr := uService.RegisterUser(*registerUserRequest)
		if rErr != nil{
			return response, fmt.Errorf("error in service register-user-request: %s", rErr.Error())

		}

		return response, nil
	case "login":
	case "whoami":
	}

	return responseParam.Response{}, nil
}
