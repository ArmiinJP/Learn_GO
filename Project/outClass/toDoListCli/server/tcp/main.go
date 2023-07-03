package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	//"os"
	"strings"

	"todolist/contract"
	"todolist/delivery/requestParam"
	"todolist/delivery/responseParam"
	"todolist/entity"
	"todolist/repository/filestorage"
	"todolist/repository/memorystorage/categorystorage"
	"todolist/repository/memorystorage/loggeduserstorage"
	"todolist/repository/memorystorage/taskstorage"
	"todolist/repository/memorystorage/userstorage"
	"todolist/service/categoryservice"
	"todolist/service/loggeduserservice"
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
	var LoggedRepo = loggeduserstorage.New()
	var categoryRepo = categorystorage.New()
	var taskRepo = taskstorage.New()


	var userService = userservice.New(&userRepo)
	var categoryService = categoryservice.New(&categoryRepo)
	var taskService = taskservice.New(&taskRepo)
	var loggedUserService = loggeduserservice.New(&LoggedRepo)

	NetworkLayer(*userAddressflag, taskService, categoryService, userService, loggedUserService)
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

func NetworkLayer(ListeningAddressPort string, tService taskservice.Service, cService categoryservice.Service, uService userservice.Service, lService loggeduserservice.Service) {
	listener, lErr := net.Listen("tcp", ListeningAddressPort)
	if lErr != nil {
		log.Fatalln("listening to the port refused: ", lErr.Error())
	}
	defer listener.Close()

	for {

		//CheckExpiresTime(lService)
		

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

		var request = &requestParam.Request{}
		uErr := json.Unmarshal(req[:numberOfByte], request)
		if uErr != nil {
			log.Printf("Unmarshaling request error : %s", uErr.Error())

			continue
		}
		request.RemoteAddr = conn.RemoteAddr().String()

		response, pErr := processRequest(*request, tService, cService, uService, lService)
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

		fmt.Println("|||||||||||||||| new ||||||||||||||||||")
		uService.Print()
		fmt.Println("----------------------------------")
		lService.Print()
		fmt.Println("----------------------------------")
		tService.Print()
		fmt.Println("----------------------------------")
		cService.Print()
		fmt.Println("|||||||||||||||||||||||||||||||||||||||")
	}
}

func processRequest(req requestParam.Request, tService taskservice.Service, cService categoryservice.Service, uService userservice.Service, lService loggeduserservice.Service) (responseParam.Response, error) {
	
	var connectionUserID int

	if !(req.Command == "register-user" || req.Command == "login"){
		response, cErr := lService.CheckLoggedINUser(requestParam.ValuesCheckedLoggedInUser{
			RemoteAddr: req.RemoteAddr,
			Time: time.Now(),
		})
		if cErr != nil{
			return response, fmt.Errorf("user with address %s not login to system. error %s", req.RemoteAddr, cErr.Error())
		}

		tmpUserID, response, rErr := lService.ReturnLoggedInUserID(requestParam.ValuesReturnLoggedInUserID{
			RemoteAddr: req.RemoteAddr,
		})
		if rErr != nil{
			return response, fmt.Errorf("%s", rErr.Error())
		}

		connectionUserID = tmpUserID
	}

	switch req.Command {
	case "create-task":
		createTaskRequestParam := &requestParam.ValuesCreateTask{}
		// createTaskRequestParam.UserID = 
		uErr := json.Unmarshal(req.ValueCommand, createTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter Create Task"}, fmt.Errorf("error unmarshaling value-create-request: %s", uErr.Error())
		}

		createTaskRequestParam.UserID = connectionUserID
		response, cErr := cService.CheckCategoryID(createTaskRequestParam.UserID, createTaskRequestParam.CategoryID)
		if cErr != nil {
			return response, fmt.Errorf(cErr.Error())
		}

		response, cErr = tService.CreateTaskRequest(*createTaskRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service create-task-request: %s", cErr.Error())
		}

		return response, nil
	case "list-task":
		listTaskRequestParam := &requestParam.ValuesListTask{}
		uErr := json.Unmarshal(req.ValueCommand, listTaskRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter List Task"}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		listTaskRequestParam.UserID = connectionUserID
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

		listTodayTaskRequestParam.UserID  = connectionUserID
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

		listSpecificDayTaskRequestParam.UserID = connectionUserID
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

		editTaskRequestParam.UserID  = connectionUserID
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
		changeStatusTaskReuquestParam := &requestParam.ValuesChangeStatusTask{UserID: connectionUserID}
		uErr := json.Unmarshal(req.ValueCommand, changeStatusTaskReuquestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter change status Task"}, fmt.Errorf("error unmarshaling value-change-status-request: %s", uErr.Error())
		}

		changeStatusTaskReuquestParam.UserID = connectionUserID
		response, cErr := tService.ChangeStatusRequest(*changeStatusTaskReuquestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service change-status-task-request: %s", cErr.Error())
		}

		return response, nil
	case "create-category":
		createCategoryRequestParam := &requestParam.ValuesCreateCategory{UserID: connectionUserID}
		uErr := json.Unmarshal(req.ValueCommand, createCategoryRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter create Category"}, fmt.Errorf("error unmarshaling value-create-category-request: %s", uErr.Error())
		}

		createCategoryRequestParam.UserID = connectionUserID
		response, cErr := cService.CreateCategoryRequest(*createCategoryRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service create-category-request: %s", cErr.Error())
		}

		return response, nil
	case "list-category":
		listCategoryRequestParam := &requestParam.ValuesListCategory{UserID: connectionUserID}
		uErr := json.Unmarshal(req.ValueCommand, listCategoryRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter List Category"}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		listCategoryRequestParam.UserID  = connectionUserID
		response, cErr := cService.ListCategoryRequest(requestParam.ValuesListCategory(*listCategoryRequestParam))
		if cErr != nil {

			return response, fmt.Errorf("error in service list-task-request: %s", cErr.Error())
		}

		return response, nil
	case "edit-category":
		editCategoryRequestParam := &requestParam.ValuesEditCategory{UserID: connectionUserID}
		uErr := json.Unmarshal(req.ValueCommand, editCategoryRequestParam)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter edit Category"}, fmt.Errorf("error unmarshaling value-edit-category-request: %s", uErr.Error())
		}

		editCategoryRequestParam.UserID = connectionUserID
		response, cErr := cService.EditCategoryRequst(*editCategoryRequestParam)
		if cErr != nil {

			return response, fmt.Errorf("error in service edit-category-request: %s", cErr.Error())
		}

		return response, nil
	case "register-user":
		registerUserRequest := &requestParam.ValuesRegisterUser{}
		uErr := json.Unmarshal(req.ValueCommand, registerUserRequest)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter Register User"}, fmt.Errorf("error unmarshaling register-user-request: %s", uErr.Error())
		}

		response, rErr := uService.RegisterUser(*registerUserRequest)
		if rErr != nil {
			return response, fmt.Errorf("error in service register-user-request: %s", rErr.Error())

		}

		return response, nil
	case "login":
		loginUserRequest := &requestParam.ValuesLoginUser{}
		uErr := json.Unmarshal(req.ValueCommand, loginUserRequest)
		if uErr != nil {

			return responseParam.Response{StatusCode: 500, Message: "Failed to Process Parameter Login User"}, fmt.Errorf("error unmarshaling login-user-request: %s", uErr.Error())
		}

		loginUserRequest.RemoteAddr = req.RemoteAddr
		response, rErr := uService.LoginUser(*loginUserRequest, lService)
		if rErr != nil {
			return response, fmt.Errorf("error in service login-user-request: %s", rErr.Error())

		}

		return response, nil
	case "whoami":
	}

	return responseParam.Response{}, nil
}

func CheckExpiresTime(lService loggeduserservice.Service){
	//checking user logged in Expire Time
	userTimeExpired := lService.CheckLoggedInTimeExpired(time.Now())
	for _, user := range userTimeExpired{
		lErr := lService.LoggedOutUser(user)
		if lErr != nil{
			log.Println(lErr.Error())
		}
	}
}