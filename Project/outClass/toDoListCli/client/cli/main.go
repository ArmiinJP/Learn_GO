package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	requestParam "todolist/delivery/requestParam"
	responseParam "todolist/delivery/response"
	"todolist/entity"
)

func main() {

	fmt.Println("Welcome toDo App")
	userCommandFlag := flag.String("command", "", "enter your command")
	targetServerFlag := flag.String("target", "", "enter server IP:PORT")
	flag.Parse()

	userCommand, userTarget := parsingFlag(*userCommandFlag, *targetServerFlag)

	for {
		dataRequest, cErr := completeCommand(userCommand)
		if cErr != nil {
			fmt.Println(cErr.Error())
		} else {
			jsonResponse, sErr := socketLayer(dataRequest, userTarget)
			if sErr != nil{
				fmt.Println(sErr.Error())
			}

			response := &responseParam.Response{}
			uErr := json.Unmarshal(jsonResponse, response)
			if uErr != nil{
				fmt.Println(uErr.Error())
			}
			fmt.Println(*response)

			dataInResponse := &[]entity.Task{}
			uErr = json.Unmarshal(response.Data, dataInResponse)
			if uErr != nil{
				fmt.Println(uErr.Error())
			}		
			
			fmt.Printf("status Code is: %d, Message is: %s, Data is: %+v", response.StatusCode, response.Message, *dataInResponse)

		}

		userCommand = giveUserCommand()
	}
}

//ParsingFlag Function do parsing Flag if exist and give flag if not exist
func parsingFlag(commandFlag, targetFlag string) (string, string) {
	

	// parsing targetFlag
	for targetFlag == "" {
		fmt.Printf("Please Enter Target:(IPServer:PORTServer): ")
		fmt.Scanln(&targetFlag)
	}

	// parsing commandFlag
	for commandFlag == "" {
		commandFlag = giveUserCommand()
	}
	
	return commandFlag, targetFlag
}

func giveUserCommand() string {
	var userCommand string

	fmt.Println("\n----------------------------------------------")
	fmt.Println("--> Accessable Command After Succseefull Login is:\n01. |create-task|", "\t02. |list-task|", "\t03. |list-today-task|", "\t04. |list-day-task|", "\t05. |edit-task|",
		"\n06. |task-complete|", "\t07. |create-category|", "\t08. |list-category|", "\t09. |edit-category|", "\t10. |whoami|", "\n11. |login|", "\t\t12. |register-user|", "\t13. |exit|")
	fmt.Println("\n--> Accessable Command without login is:\n12. |register-user|", "\t13. |exit|")
	fmt.Print("\nPlease enter your command: ")
	fmt.Scanln(&userCommand)

	return userCommand
}

func completeCommand(userCommand string) ([]byte, error) {
	switch userCommand {
	case "create-task":
		var newTask = requestParam.ValuesCreateTask{}
		fmt.Println("\n---- Creating Task")

		fmt.Printf("Please enter Task Title: ")
		fmt.Scanln(&newTask.Title)

		fmt.Printf("Please enter Task DueDate: ")
		fmt.Scanln(&newTask.DueDate)

		fmt.Printf("Please enter Task Category ID: ")
		var tmpCategoryidStr string
		fmt.Scanln(&tmpCategoryidStr)
		tmpCategoryidint, aErr := strconv.Atoi(tmpCategoryidStr)
		newTask.CategoryID = tmpCategoryidint
		if aErr != nil {

			return []byte{}, fmt.Errorf("\ncategory with id: %v is invalid", tmpCategoryidStr)
		}

		values, mErr := json.Marshal(requestParam.ValuesCreateTask{Title: newTask.Title,DueDate: newTask.DueDate,CategoryID: newTask.CategoryID,UserID: 1})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling create-task-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command: "create-task",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-task":
		values, mErr := json.Marshal(requestParam.ValuesListTask{UserID: 1})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling list-task-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command: "list-task",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-today-task":
	
	case "list-day-task":
	
	case "edit-task":
		var editTask = requestParam.ValuesEditTask{}
		fmt.Println("\n---- Editing Task")

		fmt.Printf("Please enter Task ID: ")
		var tmpTaskIDStr string
		fmt.Scanln(&tmpTaskIDStr)
		tmpTaskIDInt, aErr := strconv.Atoi(tmpTaskIDStr)
		if aErr != nil {

			return []byte{}, fmt.Errorf("\ntask with id: %v is invalid", tmpTaskIDStr)
		}
		editTask.ID = tmpTaskIDInt
		
		fmt.Printf("if want Change Title, {Enter new Task Title}, else {Just Click Enter}: ")
		fmt.Scanln(&editTask.Title)

		fmt.Printf("if want Change DueDate, {Enter new Task DueDate}, else {Just Click Enter}: ")
		fmt.Scanln(&editTask.DueDate)

		fmt.Printf("if want Change Category, {Enter new Task CategoryID}, else {Just Click Enter}: ")
		var tmpCategoryidStr string
		fmt.Scanln(&tmpCategoryidStr)
		tmpCategoryidint, aErr := strconv.Atoi(tmpCategoryidStr)
		if aErr != nil {

			return []byte{}, fmt.Errorf("\ncategory with id: %v is invalid", tmpCategoryidStr)
		}
		editTask.CategoryID = tmpCategoryidint


		values, mErr := json.Marshal(requestParam.ValuesEditTask{ID: editTask.ID, Title: editTask.Title, DueDate: editTask.DueDate, CategoryID: editTask.CategoryID,IsComplete: false, UserID: 1})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling edit-task-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command: "edit-task",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "task-complete":
	
	case "create-category":
		var newCategory = requestParam.ValuesCreateCategory{}
		fmt.Println("\n---- Creating Category")

		fmt.Printf("Please enter Category Title: ")
		fmt.Scanln(&newCategory.Title)

		fmt.Printf("Please enter Category Color: ")
		fmt.Scanln(&newCategory.Color)

		//newCategory.UserID = authenticatedUser.ID
		values, mErr := json.Marshal(requestParam.ValuesCreateCategory{Title:  newCategory.Title,Color:  newCategory.Color,UserID: 1})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling create-category-value: %s", mErr.Error())
		}
		req := requestParam.Request{
			Command: "create-category",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-category":
		values, mErr := json.Marshal(requestParam.ValuesListCategory{UserID: 1})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling list-category-value: %s", mErr.Error())
		}		
		req := requestParam.Request{
			Command: "list-category",
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

		values, mErr := json.Marshal(requestParam.ValuesRegisterUser{Email:    newUser.Email,Password: newUser.Password})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling register-user-value: %s", mErr.Error())
		}

		req := requestParam.Request{
			Command: "register-User",
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


		values, mErr := json.Marshal(requestParam.ValuesLoginUser{Email:    newUser.Email,Password: newUser.Password})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling login-user-value: %s", mErr.Error())
		}
		
		req := requestParam.Request{
			Command: "login",
			ValueCommand: values,
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "whoami":
		values, mErr := json.Marshal(requestParam.ValuesWhoami{UserID: 1})
		if mErr != nil{
			return []byte{}, fmt.Errorf("error Marshaling whoami-value: %s", mErr.Error())
		}		
		req := requestParam.Request{
			Command: "whoami",
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
		fmt.Printf("\n--- command %s is not found!!\n", userCommand)
	}

	return []byte{}, nil
}

func socketLayer(data []byte, target string) ([]byte, error){
	
	conn, dErr := net.Dial("tcp", target)
	if dErr != nil{
		return []byte{}, fmt.Errorf("error Dialing: %s", dErr.Error())
	}

	_, wErr := conn.Write(data)
	if wErr != nil{
		return []byte{}, fmt.Errorf("error Sending data to Server: %s", wErr.Error())
	}

	var res = make([]byte, 1024)
	numberOfByte, rErr := conn.Read(res)
	if rErr != nil{
		return []byte{}, fmt.Errorf("error Reading data from Server: %s", rErr.Error())
	}

	return res[:numberOfByte], nil
}