package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	//"errors"
	//"log"
)

type User struct {
	ID       int
	Email    string
	Password string
}

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID   int
	IsComplete bool
	UserID     int
}

type Category struct {
	ID	   int
	Title  string
	Color  string
	UserID int
}

var Users []User
var Tasks []Task
var Categoreis []Category
var authenticatedUser *User

func LoginUser() {
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
		}
	}

	if authenticatedUser == nil{
		fmt.Println("\n----- Faild Logging !!")
	}
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
func CreateTask(){
	var newTask Task
	fmt.Println("\n---- Creating Task")

	fmt.Printf("Please enter Task Title: ")
	fmt.Scanln(&newTask.Title)	
	fmt.Printf("Please enter Task DueDate: ")
	fmt.Scanln(&newTask.DueDate)	
	
	//validating category Exist (int && category user exist)
	fmt.Printf("Please enter Task Category ID: ")
	var tmpCategoryidStr string
	fmt.Scanln(&tmpCategoryidStr)
	tmpCategoryidInt, err := strconv.Atoi(tmpCategoryidStr)
	if err != nil{
		fmt.Printf("\nCategory with id: %v is invalid!!\n", tmpCategoryidStr)
		
		return
	}

	CategoryFound := false
	for _, cat := range Categoreis{
		if cat.UserID == authenticatedUser.ID && cat.ID == tmpCategoryidInt{
			newTask.CategoryID = tmpCategoryidInt
			CategoryFound = true
			break
		}
	}
	if CategoryFound == false{
		fmt.Printf("\nCategory with id: %d is not Found!!\n", tmpCategoryidInt)
		
		return
	}

	newTask.ID = len(Tasks) + 1
	newTask.IsComplete = false
	newTask.UserID = authenticatedUser.ID
	
	Tasks = append(Tasks, newTask)
	fmt.Println("Task Successfully Added")
}
func ListTask(){
	for _, v := range Tasks{
		if v.UserID == authenticatedUser.ID {
			fmt.Println("----------\ntask name is:", v.Title,
						"\ntask category ID is:", v.CategoryID, 
						"\ntask dueDate is:", v.DueDate, 
						"\ntask completed is:", v.IsComplete)
		}
	}
}
func ListTodayTask(){}
func ListDayTask()      {}
func EditTask()         {}
func ChangeStatusTask() {}

func CreateCategory(){
	var newCategory Category
	fmt.Println("\n---- Creating Category")

	fmt.Printf("Please enter Category Title: ")
	fmt.Scanln(&newCategory.Title)	
	fmt.Printf("Please enter Category Color: ")
	fmt.Scanln(&newCategory.Color)	
	newCategory.ID = len(Categoreis) + 1
	newCategory.UserID = authenticatedUser.ID
	
	Categoreis = append(Categoreis, newCategory)
	fmt.Println("Category Successfully Added")	
}
func ListCategory(){
	for _, v := range Categoreis{
		if v.UserID == authenticatedUser.ID {
			fmt.Println("Category name is:", v.Title,
						"Category ID is:", v.ID,
						"\nCategory Color is:", v.Color)
		}
	}
}
func EditCategory(){}

func RunCommand(userCommand string){

	//service need logging before use, except exit & register-user
	if userCommand != "exit" && userCommand != "register-user" && authenticatedUser == nil{
		LoginUser()
		if authenticatedUser == nil {

			return
		}
	}

	switch userCommand {
		case "create-task":
			CreateTask()
		case "list-task":
			ListTask()
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
}
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
		RunCommand(userCommand)

		//get new command from user
		fmt.Print("\nPlease enter your command: ")
		isEmpty, _ := fmt.Scanln(&userCommand)
		if isEmpty == 0 {
			userCommand = "empty"
		}
	}
}