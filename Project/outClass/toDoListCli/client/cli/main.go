package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	requestParam "todolist/delivery/requestParam"
)

func main() {

	fmt.Println("Welcome toDo App")
	userCommandflag := flag.String("command", "", "enter your command")
	flag.Parse()

	userCommand := parsingFlag(*userCommandflag)

	// task service running

	for {
		dataRequest, cErr := completeCommand(userCommand /* inja task service*/)
		if cErr != nil {
			fmt.Println(cErr.Error())
		} else {
			//networking
			// test := parameterCommand.CreateTask{}
			// json.Unmarshal(dataRequest, &test)
			// fmt.Println(test)
			fmt.Println(string(dataRequest))
		}
		userCommand = giveUserCommand()
	}
}

func parsingFlag(commandFlag string) string {

	//This function parsing Flag if exist and give flag if not exist
	var userCommand string

	// parsing commandFlag
	if commandFlag == "" {
		commandFlag = giveUserCommand()
	}
	userCommand = commandFlag

	return userCommand
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

		//working this line:
		//newTask.UserID = 1

		req := requestParam.CreateTask{
			Command: "create-task",
			ValueCommand: requestParam.ValuesCreateTask{
				Title:      newTask.Title,
				DueDate:    newTask.DueDate,
				CategoryID: newTask.CategoryID,
				UserID:     1,
			},
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-task":
		req := requestParam.ListTask{
			Command: "list-task",
			ValueCommand: requestParam.ValuesListTask{
				UserID: 1,
			},
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-today-task":
	case "list-day-task":
	case "edit-task":
	case "task-complete":
	case "create-category":
		var newCategory = requestParam.ValuesCreateCategory{}
		fmt.Println("\n---- Creating Category")

		fmt.Printf("Please enter Category Title: ")
		fmt.Scanln(&newCategory.Title)

		fmt.Printf("Please enter Category Color: ")
		fmt.Scanln(&newCategory.Color)

		//newCategory.UserID = authenticatedUser.ID

		req := requestParam.CreateCategory{
			Command: "create-category",
			ValueCommand: requestParam.ValuesCreateCategory{
				Title:  newCategory.Title,
				Color:  newCategory.Color,
				UserID: 1,
			},
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "list-category":
		req := requestParam.ListCategory{
			Command: "list-category",
			ValueCommand: requestParam.ValuesListCategory{
				UserID: 1,
			},
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

		req := requestParam.RegisterUser{
			Command: "register-User",
			ValueCommand: requestParam.ValuesRegisterUser{
				Email:    newUser.Email,
				Password: newUser.Password,
			},
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

		req := requestParam.LoginUser{
			Command: "login",
			ValueCommand: requestParam.ValuesLoginUser{
				Email:    newUser.Email,
				Password: newUser.Password,
			},
		}

		dataRequest, jErr := json.Marshal(&req)
		if jErr != nil {
			return []byte{}, fmt.Errorf("error in Marshaling data: %s", jErr.Error())
		}

		return dataRequest, nil

	case "whoami":
		req := requestParam.Whoami{
			Command: "whoami",
			ValueCommand: requestParam.ValuesWhoami{
				UserID: 1,
			},
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
