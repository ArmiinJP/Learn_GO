package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"todolistapp/delivery/param"
)

func main(){

	command := flag.String("command", "empty", "")
	address := flag.String("address", "empty", "")
	flag.Parse()
	if *address == "empty"{
		fmt.Println("address is invalid")

		return
	}
	
	//client start connect to server
	connection, cErr := net.Dial("tcp", *address)
	if cErr != nil{
		log.Fatalln("error connection")
	}


	writeData := []byte{}
	switch *command {
		case "create-task":
			req := param.Request{
				Command: "create-task", 
				Values: param.RequestValue{
					CreatedTaskRequest: param.CreatedTaskRequest{Title: "test",DueDate: "test",CategoryID: 1},},
			}

			data, mErr := json.Marshal(&req)
			if mErr != nil{
				log.Fatalln("error :", mErr)
			}
			writeData = data
		}

		// conncetion is ready and client send data to server
		_, wErr := connection.Write(writeData)
		if wErr != nil {
			log.Fatalln("error when writeing buffer", wErr)
		}
		
		// client stop to recive reponse from server that data succefully recive from server
		var response = make([]byte, 1024)
		if _, rErr := connection.Read(response); rErr != nil{
			log.Fatalln("error when getting response")
		}
		
		//Print all event in terminal
		fmt.Printf("response is: %s\n", string(response))
}