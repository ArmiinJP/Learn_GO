package main

import (
	"encoding/json"
	"encoding/xml"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"log"
	"crypto/sha512"


	//"errors"
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

	usersFileFormat string
)

func main() {
	
	fmt.Println("Welcome toDo App")
	userCommandflag := flag.String("command", "", "enter your command")
	userFormatFlag := flag.String("format", "", "enter your format to save file")
	flag.Parse()

	datasetNameExist := checkFormatFile(*userFormatFlag)

	// reading dataset if dataset exist
	if datasetNameExist != "" {
		err := readDataset(datasetNameExist)
		if err != nil{

			log.Fatalln(err)
		}
	}

	userCommand := *userCommandflag

	for {
		userCommand = giveUserCommand(userCommand)
		RunCommand(userCommand)
		userCommand = ""
	}
}

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
	newUser.Password = hashPassword([]byte(newUser.Password))

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
	newUser.Password = hashPassword([]byte(newUser.Password))
	//isEmpty, _ = fmt.Scanln(&newUser.Password) //check input

	newUser.ID = len(Users) + 1

	if err := saveToFile(newUser); err != nil{
		fmt.Printf("\nRegister user Failed!!\n")	
		return
	}
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
	if !CategoryFound{
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
						"\nCategory ID is:", v.ID,
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

func checkFormatFile(userFormatFlag string) string{
	//datase --> if exist 	  : return name dataset and set usersFileFormat to format dataset
	//		 	 if not exist : return "" and set usersFileFormat to user flag format or default format(json)
	var fileFormats = []string{"data.csv", "data.json", "data.xml", "data.txt"}
	var err error
	var fileDatasetName string

	for _, v := range fileFormats{
		_ , err = os.Stat(v)
		if !os.IsNotExist(err) {
			fileDatasetName = v
			break
		}
	}
	// datase not exist
	if fileDatasetName == ""{
		usersFileFormat = strings.ToLower(userFormatFlag)
		
		switch usersFileFormat{
			case "":
				fmt.Println("Format not Determine\n so default format (json) is apply")
				usersFileFormat = "json"
			
			case "json", "xml", "csv", "txt":
				//usersFileFormat = usersFileFormat
			
			default:
				fmt.Println("Format File you choose is False\n so default format (json) is apply")
				usersFileFormat = "json"
		}
	
		return ""

	// datase exist
	} else {
		usersFileFormat = strings.Split(fileDatasetName, ".")[1]
		fmt.Printf("file with format %s is Exist!\nso Data is save with format: | %s |\n",
					 fileDatasetName,strings.ToUpper(strings.Split(fileDatasetName, ".")[1]))
		
		return fileDatasetName
	}
}

func readDataset(datasetName string) error{
    // open file
    f, err := os.Open(datasetName)
    if err != nil {
        return fmt.Errorf("error when open Dataset")
    }
    defer f.Close()

	var data string
	var fetchUser User
    scanner := bufio.NewScanner(f)
    
	for scanner.Scan() {
		var err error
		// delete "\n" of end of each line
        data = strings.Split(scanner.Text(), "\n")[0]

		switch usersFileFormat {
		case "json": 
			err = json.Unmarshal([]byte(data), &fetchUser)
			if err != nil{

				return fmt.Errorf("datas in dataset is wrong format for JSON")
			}
		
		case "xml":
			err = xml.Unmarshal([]byte(data), &fetchUser)
			if err != nil{

				return fmt.Errorf("datas in dataset is wrong format for JSON")
			}			
		
		case "csv":
			fetchUser.Email = strings.Split(data, ",")[0]
			fetchUser.ID, err = strconv.Atoi(strings.Split(data, ",")[1])
			if err != nil{

				return fmt.Errorf("datas in dataset is wrong format for CSV")
			}
			fetchUser.Password = strings.Split(data, ",")[2]
		
		case "txt":
			fetchUser.ID, err = strconv.Atoi(strings.Split(strings.Split(data, ",")[0], ": ")[1])
			if err != nil{
				return fmt.Errorf("datas in dataset is wrong format for TXT")

			}
			fetchUser.Email = strings.Split(strings.Split(data, ",")[1], ": ")[1]
			fetchUser.Password = strings.Split(strings.Split(data, ",")[2], ": ")[1]
		}

		Users = append(Users, fetchUser)
    }

	return nil
}

func saveToFile(newUser User) error{
	//Choose format to save into file
	var data []byte
	var err error
	var nameFile string

	switch usersFileFormat {
	case "json":
		nameFile = "data.json"
		data, err = json.Marshal(&newUser)
		if err != nil {

			return fmt.Errorf("could not convert data to json")
		}		

	case "xml":
		nameFile = "data.xml"
		data, err = xml.Marshal(&newUser)
		if err != nil {

			return fmt.Errorf("could not convert data to xml")
		}

	case "csv":
		nameFile = "data.csv"
		data = []byte(fmt.Sprintf("%s,%d,%s", newUser.Email, newUser.ID, newUser.Password))
	
	case "txt":
		nameFile = "data.txt"
		data = []byte(fmt.Sprintf("id: %d, email: %s, password: %s", newUser.ID, newUser.Email, newUser.Password))
	
	default:

		return fmt.Errorf("format to save file is invalid")
	}

	//write data to the file
    file, err := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    
    if err != nil {
    	
		return fmt.Errorf("could not open dataset")
	}

	defer file.Close()
	 
	_, err2 := file.WriteString(string(data)+"\n")

	if err2 != nil {

		return fmt.Errorf("could not write data to dataset")
	}

	return nil
}

func hashPassword(password []byte ) string{
	hash := sha512.New()
	hash.Write(password)

	encodedHash := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return encodedHash
}