package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"todolist/contract"
	"todolist/delivery/requestParam"
	"todolist/entity"
	"todolist/repository/filestorage"
	"todolist/repository/memorystorage/categorystorage"
	"todolist/repository/memorystorage/taskstorage"
	categoryservice "todolist/service/category"
	taskservice "todolist/service/task"
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

	var taskRepo = taskstorage.New()
	var CategoryRepo = categorystorage.New()
	
	var categoryService = categoryservice.New(&CategoryRepo)
	var taskService = taskservice.New(&taskRepo)

	// netwroking
	listener, lErr := net.Listen("tcp", *userAddressflag)
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

		dataResponse, pErr := processRequest(request, taskService, categoryService)
		if pErr != nil {
			log.Println("Error Processing request: ", pErr.Error())

			continue
		}

		_, wErr := conn.Write(dataResponse)
		if wErr != nil {
			log.Println("Error responsing Client: ", wErr.Error())

			continue
		}

		taskRepo.Print()
	}

}

func processRequest(req requestParam.Request, taskService taskservice.Service, categoryService categoryservice.Service) ([]byte, error) {
	switch req.Command {
	case "create-task":
		createTaskRequestParam := &requestParam.ValuesCreateTask{}
		uErr := json.Unmarshal(req.ValueCommand, createTaskRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-create-request: %s", uErr.Error())
		}

		response, cErr := taskService.CreateTaskRequest(*createTaskRequestParam)
		
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service create-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "list-task":
		listTaskRequestParam := &requestParam.ValuesListTask{}
		uErr := json.Unmarshal(req.ValueCommand, listTaskRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		response, cErr := taskService.ListTaskRequest(*listTaskRequestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service list-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "list-today-task":
		listTodayTaskRequestParam := &requestParam.ValueslistTodayTask{}
		uErr := json.Unmarshal(req.ValueCommand, listTodayTaskRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-list-today-request: %s", uErr.Error())
		}

		response, cErr := taskService.ListTodayRequest(*listTodayTaskRequestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service list-today-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "list-day-task":
		listSpecificDayTaskRequestParam := &requestParam.ValuesListSpecificDayTask{}
		uErr := json.Unmarshal(req.ValueCommand, listSpecificDayTaskRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-list-specific-day-request: %s", uErr.Error())
		}

		response, cErr := taskService.ListSpecificDayRequest(*listSpecificDayTaskRequestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service list-specific-day-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "edit-task":
		editTaskRequestParam := &requestParam.ValuesEditTask{}
		uErr := json.Unmarshal(req.ValueCommand, editTaskRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-edit-request: %s", uErr.Error())
		}

		response, cErr := taskService.EditTaskRequst(*editTaskRequestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service edit-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "change-status-task":
		changeStatusTaskReuquestParam := &requestParam.ValuesChangeStatusTask{}
		uErr := json.Unmarshal(req.ValueCommand, changeStatusTaskReuquestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-change-status-request: %s", uErr.Error())
		}

		response, cErr := taskService.ChangeStatusRequest(*changeStatusTaskReuquestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service change-status-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "create-category":
		createCategoryRequestParam := &requestParam.ValuesCreateCategory{}
		uErr := json.Unmarshal(req.ValueCommand, createCategoryRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-create-category-request: %s", uErr.Error())
		}

		response, cErr := categoryService.CreateCategoryRequest(*createCategoryRequestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service create-category-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "list-category":
		listCategoryRequestParam := &requestParam.ValuesListCategory{}
		uErr := json.Unmarshal(req.ValueCommand, listCategoryRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		response, cErr := categoryService.ListCategoryRequest(requestParam.ValuesListCategory(*listCategoryRequestParam))
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service list-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "edit-category":
		editCategoryRequestParam := &requestParam.ValuesEditCategory{}
		uErr := json.Unmarshal(req.ValueCommand, editCategoryRequestParam)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-edit-category-request: %s", uErr.Error())
		}

		response, cErr := categoryService.EditCategoryRequst(*editCategoryRequestParam)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service edit-category-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "register-user":
	case "login":
	case "whoami":
	case "exit":
		fmt.Println("App is Closed")
		os.Exit(0)
	default:
		fmt.Printf("\n--- command %s is not found!!\n", req.Command)
	}

	return []byte{}, nil
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
