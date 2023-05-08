package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
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
			
		
		var request = requestParam.CreateTask{}
		json.Unmarshal(req[:numberOfByte], &request)

		response, _ := taskService.CreateTaskRequest(request.ValueCommand)
		res, _ := json.Marshal(response)

		conn.Write(res)

		taskRepo.Print()
	}

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
