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
	"todolist/repository/memorystorage"
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

	var taskRepo = memorystorage.TaskStorage{}
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

		dataResponse, pErr := processRequest(request, taskService)
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

func processRequest(req requestParam.Request, taskService taskservice.Service) ([]byte, error) {
	switch req.Command {
	case "create-task":
		createTaskRequest := &requestParam.ValuesCreateTask{}
		uErr := json.Unmarshal(req.ValueCommand, createTaskRequest)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-create-request: %s", uErr.Error())
		}

		response, cErr := taskService.CreateTaskRequest(*createTaskRequest)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service create-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil
	case "list-task":
		listTaskRequest := &requestParam.ValuesListTask{}
		uErr := json.Unmarshal(req.ValueCommand, listTaskRequest)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-list-request: %s", uErr.Error())
		}

		response, cErr := taskService.ListTaskRequest(*listTaskRequest)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service list-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil

	case "list-today-task":
	case "list-day-task":
	case "edit-task":
		editTaskRequest := &requestParam.ValuesEditTask{}
		uErr := json.Unmarshal(req.ValueCommand, editTaskRequest)
		if uErr != nil {

			return []byte{}, fmt.Errorf("error unmarshaling value-edit-request: %s", uErr.Error())
		}

		response, cErr := taskService.EditTaskRequst(*editTaskRequest)
		if cErr != nil {

			return []byte{}, fmt.Errorf("error in service edit-task-request: %s", cErr.Error())
		}

		dataResponse, mErr := json.Marshal(response)
		if mErr != nil {

			return []byte{}, fmt.Errorf("error Marshaling response %s", mErr.Error())
		}

		return dataResponse, nil

	case "task-complete":
	case "create-category":
		var newCategory = requestParam.ValuesCreateCategory{}
		fmt.Println("\n---- Creating Category")

		fmt.Printf("Please enter Category Title: ")
		fmt.Scanln(&newCategory.Title)

		fmt.Printf("Please enter Category Color: ")
		fmt.Scanln(&newCategory.Color)

		//newCategory.UserID = authenticatedUser.ID
		values, mErr := json.Marshal(requestParam.ValuesCreateCategory{Title: newCategory.Title, Color: newCategory.Color, UserID: 1})
		if mErr != nil {
			return []byte{}, fmt.Errorf("error Marshaling create-category-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command:      "create-category",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-category":
		values, mErr := json.Marshal(requestParam.ValuesListCategory{UserID: 1})
		if mErr != nil {
			return []byte{}, fmt.Errorf("error Marshaling list-category-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command:      "list-category",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "edit-category":
	case "register-user":
		newUser := requestParam.ValuesRegisterUser{}
		fmt.Println("\n----- Registering User ----- ")

		fmt.Printf("Please enter your Email: ")
		fmt.Scanln(&newUser.Email)
		//isEmpty, _ = fmt.Scanln(&newUser.Email) //check input

		fmt.Printf("Please enter your Password: ")
		fmt.Scanln(&newUser.Password)
		//isEmpty, _ = fmt.Scanln(&newUser.Password) //check input

		values, mErr := json.Marshal(requestParam.ValuesRegisterUser{Email: newUser.Email, Password: newUser.Password})
		if mErr != nil {
			return []byte{}, fmt.Errorf("error Marshaling register-user-value: %s", mErr.Error())
		}

		req := requestParam.Request{
			Command:      "register-User",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "login":
		newUser := requestParam.ValuesLoginUser{}
		fmt.Println("\n----- Logging User ----- ")

		fmt.Printf("Please enter your Email: ")
		fmt.Scanln(&newUser.Email)
		//isEmpty, _ = fmt.Scanln(&newUser.Email) //check input

		fmt.Printf("Please enter your Password: ")
		fmt.Scanln(&newUser.Password)
		//isEmpty, _ = fmt.Scanln(&newUser.Password) //check input

		values, mErr := json.Marshal(requestParam.ValuesLoginUser{Email: newUser.Email, Password: newUser.Password})
		if mErr != nil {
			return []byte{}, fmt.Errorf("error Marshaling login-user-value: %s", mErr.Error())
		}

		req := requestParam.Request{
			Command:      "login",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "whoami":
		values, mErr := json.Marshal(requestParam.ValuesWhoami{UserID: 1})
		if mErr != nil {
			return []byte{}, fmt.Errorf("error Marshaling whoami-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command:      "whoami",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

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
