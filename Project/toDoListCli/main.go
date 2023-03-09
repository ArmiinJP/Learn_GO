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

var (
	Users []User
	Tasks []Task
	Categoreis []Category

	authenticatedUser *User
)

func LoginUser() {
	//var isEmpty int
	fmt.Println("\n----- Logging User ----- ")
	var newUser User
	var userInput string

	if authenticatedUser != nil {
		fmt.Printf("User %s is now logged in\nIs User Logged out?(Y/N): ", authenticatedUser.Email)
		fmt.Scanln(&userInput)
		
		switch userInput{
		case "y", "Y", "yes", "Yes", "YES":
			fmt.Printf("User %s Is Successfully Logged out\n\n", authenticatedUser.Email)
			authenticatedUser = nil
		case "n", "N", "no", "NO", "No":
			fmt.Printf("Login new User Stop, User %s still Logged in\n", authenticatedUser.Email)

			return

		default:
			fmt.Println("your Input is False")
			
			return
		}
	}


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
	if userCommand != "exit" && userCommand != "register-user" && userCommand != "" && authenticatedUser == nil {
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
		case "":
			fmt.Printf("\n--- command not input!!\n")
		case "exit":
			fmt.Println("App is Closed")
			os.Exit(0)					
		default:
			fmt.Printf("\n--- command %s is not found!!\n", userCommand)
	}
}
func giveUserCommand(userFlag string) string{
	var userCommand string
	if userFlag == ""{
		if authenticatedUser == nil{
			fmt.Println("\n-----------------------User not Login the APP-----------------------")
			fmt.Println("--> Accessable Command After Succseefull Login is:\n01. |create-task|", "\t02. |list-task|", "\t03. |list-today-task|", "\t04. |list-day-task|", "\t05. |edit-task|",
						"\n06. |task-complete|", "\t07. |create-category|", "\t08. |list-category|", "\t09. |edit-category|", "\t10. |whoami|", "\n11. |login|", "\t\t12. |register-user|", "\t13. |exit|")
			fmt.Println("\n--> Accessable Command without login is:\n12. |register-user|", "\t13. |exit|")
		} else {
			fmt.Println("\n---------------------------User Logged in---------------------------")
			fmt.Println("--> Accessable Command is:\n01. |create-task|", "\t02. |list-task|", "\t03. |list-today-task|", "\t04. |list-day-task|", "\t05. |edit-task|",
						"\n06. |task-complete|", "\t07. |create-category|", "\t08. |list-category|", "\t09. |edit-category|", "\t10. |whoami|", "\n11. |login|", "\t\t12. |register-user|", "\t13. |exit|")		}

		fmt.Print("\nPlease enter your command: ")
		fmt.Scanln(&userCommand)
	} else{
		userCommand = userFlag
	}
	
	return userCommand
}
func main() {
	fmt.Println("Welcome toDo App")
	userCommandflag := flag.String("command", "", "enter your command")
	flag.Parse()

	userCommand := *userCommandflag
	for {
		userCommand = giveUserCommand(userCommand)
		RunCommand(userCommand)
		userCommand = ""		
	}
}