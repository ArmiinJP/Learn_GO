package main

import (
	"flag"
	"fmt"
	"os"
	"errors"
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
var Tasks []Task
var authenticatedUser *User

func LoginUser() error {
	var newUser User
	//var isEmpty int
	fmt.Println("\n----- Logging User ----- ")

	fmt.Printf("Please enter your Email: ")
	fmt.Scanln(&newUser.Email)
	//isEmpty, _ = fmt.Scanln(&newUser.Email) //check input

	fmt.Printf("Please enter your Password: ")
	fmt.Scanln(&newUser.Password)

	for _, user := range Users {
		if user.Email == newUser.Email && user.Password == newUser.Password {
			newUser.ID = user.ID
			authenticatedUser = &newUser
			fmt.Println("\n----- Successfull Logging")
			
			return nil
		}
	}

	if authenticatedUser == nil{
		
		return errors.New("\n----- Faild Logging !!")
	}
	
	return nil
}
func RegisterUser() {
	var newUser User
	//var isEmpty int
	fmt.Println("\n----- Registering User ----- ")

	fmt.Printf("Please enter your Email: ")
	fmt.Scanln(&newUser.Email)
	//isEmpty, _ = fmt.Scanln(&newUser.Email) //check input

	fmt.Printf("Please enter your Password: ")
	fmt.Scanln(&newUser.Password)
	//isEmpty, _ = fmt.Scanln(&newUser.Password) //check input

	newUser.ID = len(Users) + 1
	Users = append(Users, newUser)
	fmt.Printf("\nUser with Email: %s\n---> Successfull Registerd\n", newUser.Email)
}
func CreateTask(id int){
	var newTask Task
	fmt.Println("\n---- Creating Task")

	fmt.Printf("Please enter Task Title: ")
	fmt.Scanln(&newTask.Title)	
	fmt.Printf("Please enter Task DueDate: ")
	fmt.Scanln(&newTask.DueDate)	
	fmt.Printf("Please enter Task Category: ")
	fmt.Scanln(&newTask.Category)	
	newTask.IsComplete = false
	newTask.UserID = id
	
	Tasks = append(Tasks, newTask)
	fmt.Println("Task Successfully Added")
}
func ListTodayTask()    {}
func ListDayTask()      {}
func EditTask()         {}
func ChangeStatusTask() {}

func CreateCategory() {}
func ListCategory()   {}
func EditCategory()   {}

func main() {
	//fmt.Println("test")
	userCommandflag := flag.String("command", "empty", "enter your command")
	flag.Parse()

	// if flag command not input, ask from user to input command
	for *userCommandflag == "empty"{
		fmt.Print("Please enter your command: ")
		isEmpty, _ := fmt.Scanln(userCommandflag)
		if isEmpty == 0 {
			*userCommandflag = "empty"
		}
	}

	// loop{} to run app when user input "exit" command
	userCommand := *userCommandflag
	for {
		
		//service need logging before use, except exit & register-user
		if userCommand != "exit" && userCommand != "register-user" {
			if authenticatedUser == nil {
				err := LoginUser()
				if err != nil {
					fmt.Println(err)

					continue
				}
			}
		}
	
		switch userCommand {
			case "create-task":
				CreateTask(authenticatedUser.ID)
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
			case "register-user":
				RegisterUser()
			case "login":
				LoginUser()
			case "whoami":
				fmt.Printf("\n you're ID is: %d, and you're Email is: %s\n", authenticatedUser.ID, authenticatedUser.Email)
			case "empty":
				fmt.Printf("\n--- command not input!!\n")
			case "exit":
				os.Exit(0)					
			default:
				fmt.Printf("\n--- command %s is not found!!\n", userCommand)
		}
	
		//get new command from user
		fmt.Print("\nPlease enter your command: ")
		isEmpty, _ := fmt.Scanln(&userCommand)
		if isEmpty == 0 {
			userCommand = "empty"
		}
	}
}

