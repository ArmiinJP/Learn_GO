package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	//"log"
)

type User struct {
	ID       int
	Email    string
	Password string
}

type Task struct {
	Title      string
	DueDate    string
	Category   string //question
	IsComplete bool
	UserID     int
}

type Category struct {
	Title  string
	Color  string
	UserID int
}

var Users []User
var authenticatedUser *User

func LoginUser() {
	var newUser User
	//var isEmpty int
	fmt.Println("----- Logging User ----- ")

	fmt.Printf("Please enter your Email: ")
	fmt.Scanln(&newUser.Email)
	//isEmpty, _ = fmt.Scanln(&newUser.Email) //check input

	fmt.Printf("Please enter your Password: ")
	fmt.Scanln(&newUser.Password)

	for _, user := range Users {
		if user.Email == newUser.Email && user.Password == newUser.Password {
			authenticatedUser = &newUser
		}
	}
}
func RegisterUser() {
	var newUser User
	//var isEmpty int
	fmt.Println("----- Registering User ----- ")

	fmt.Printf("Please enter your Email: ")
	fmt.Scanln(&newUser.Email)
	//isEmpty, _ = fmt.Scanln(&newUser.Email) //check input

	fmt.Printf("Please enter your Password: ")
	fmt.Scanln(&newUser.Password)
	//isEmpty, _ = fmt.Scanln(&newUser.Password) //check input

	newUser.ID = len(Users) + 1
	Users = append(Users, newUser)
}

func CreateTask()       {}
func ListTodayTask()    {}
func ListDayTask()      {}
func EditTask()         {}
func ChangeStatusTask() {}

func CreateCategory() {}
func ListCategory()   {}
func EditCategory()   {}

func main() {
	//fmt.Println("test")
	userCommand := flag.String("command", "empty", "enter your command")
	flag.Parse()

	for {

		if *userCommand != "exit" && *userCommand != "register-user" && *userCommand != "empty" {
			LoginUser()
		} else {
			switch *userCommand{
			case "register-user":
				RegisterUser()
			case "exit":
				os.Exit(0)
			}
		}

		if authenticatedUser != nil{
			switch *userCommand {
				case "create-task":
					CreateTask()
				case "list-today-task":
					ListTodayTask()
				case "list-day-task":
					ListDayTask()
				case "edit-task":
					EditTask()
				case "task-complete":
					ChangeStatusTask()
				case "create-category":
					CreateCategory()
				case "list-category":
					ListCategory()
				case "edit-category":
					EditCategory()
				case "empty":
					fmt.Printf("--- command not input!!\n")
				default:
					fmt.Printf("--- command %s is not found!!\n", *userCommand)
			}

		fmt.Print("Please enter your command: ")
		isEmpty, _ := fmt.Scanln(userCommand)
		if isEmpty == 0 {
			*userCommand = "empty"
		}

	}
}
