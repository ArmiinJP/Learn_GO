package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"

	// "encoding/json"
	// "errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	// "strings"

	"todolistapp/constant"
	"todolistapp/contract"
	"todolistapp/entity"
	"todolistapp/service/task"

	// "todolist/fileStore2"
	"todolistapp/repository/filestore"
	"todolistapp/repository/memorystore"
)


var (
	userStorage     []entity.User
	categoryStorage []entity.Category
	
	authenticatedUser *entity.User
	serializationMode string

)


func  main() {

	taskMemoryRepo := memorystore.NewTask()

	taskService := task.NewService(taskMemoryRepo)

	serializeMode := flag.String("serialize-mode", constant.ManDarAvardiSerializationMode, "serialization mode to write data to file")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()
	fmt.Println("Hello to TODO app")




	switch *serializeMode {
	case constant.ManDarAvardiSerializationMode:
		serializationMode = constant.ManDarAvardiSerializationMode
	default:
		serializationMode = constant.JsonSerializationMode
	}

	// userfileStore := filestore2.FileStorage2{}
	// var ur contract.UserStoreRead = userfileStore

	// just this file change
	userfileStore := filestore.New("user.txt", serializationMode)
	var ur contract.UserStoreRead = userfileStore

	userStorage = append(userStorage, ur.Read()...) 


	

	for {
		runCommand(*command, userfileStore, &taskService)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string, ui contract.UserStoreWrite, taskService *task.Service) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}

	
	switch command {
	case "create-task":
		createTask(taskService)
	case "create-category":
		createCategory()
	case "register-user":
		registerUser(ui)
	case "list-task":
		listTask(taskService)
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask(taskService *task.Service) {

	scanner := bufio.NewScanner(os.Stdin)
	var title, duedate, category string

	fmt.Println("please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category id")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category-id is not valid integer, %v\n", err)

		return
	}

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	task, err := taskService.CreatedTask(task.CreatedRequest{
		Title: title,
		DueDate: duedate,
		CategoryID: categoryID,
		AutheticatedUserID: authenticatedUser.ID,
	})
	if err != nil{
		fmt.Println("error", err)
	}

	fmt.Println(task.Task)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the category color")
	scanner.Scan()
	color = scanner.Text()
	fmt.Println("category", title, color)

	c := entity.Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, c)
}

func registerUser(ui contract.UserStoreWrite) {
	scanner := bufio.NewScanner(os.Stdin)
	var id, name, email, password string

	fmt.Println("please enter the name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("please enter the email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user:", id, email, password)

	user := entity.User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashThePassword(password),
	}

	userStorage = append(userStorage, user)

	ui.Save(user)
}

func login() {
	fmt.Println("login process")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == hashThePassword(password) {
			authenticatedUser = &user
			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct")
	}
}

func listTask(taskService *task.Service) {
	
	userTasks, err := taskService.ListTask(task.ListRequest{UserID: authenticatedUser.ID})
	if err != nil{
		fmt.Println("error", err)
	}

	fmt.Println(userTasks.Tasks)
}

// func loadUserStorage (ur contract.UserStoreRead) {
// 	users := ur.Read()
// 	userStorage = append(userStorage, users...)
// }


func hashThePassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}