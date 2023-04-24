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

	"session_2_2/constant"
	"session_2_2/contract"
	"session_2_2/entity"
	// "session_2_2/fileStore2"
	"session_2_2/filestore"
)


type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var (
	userStorage     []entity.User
	categoryStorage []Category
	taskStorage     []Task

	authenticatedUser *entity.User
	serializationMode string

)


func main() {
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
		runCommand(*command, userfileStore)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string, ui contract.UserStoreWrite) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()

		if authenticatedUser == nil {
			return
		}
	}

	
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser(ui)
	case "list-task":
		listTask()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
	}
}

func createTask() {

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

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Printf("category-id is not found\n")

		return
	}

	fmt.Println("please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	// validation
	// category validate

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)
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

	c := Category{
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

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}

// func loadUserStorage (ur contract.UserStoreRead) {
// 	users := ur.Read()
// 	userStorage = append(userStorage, users...)
// }


func hashThePassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}

