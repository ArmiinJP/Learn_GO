package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"todolistapp/delivery/param"
)

func main(){

	//client start connect to server
	connection, cErr := net.Dial("tcp", "127.0.0.1:2023")
	if cErr != nil{
		log.Fatalln("error connection")
	}
	
	req := param.Request{Command: "create-task"}
	data, mErr := json.Marshal(&req)
	if mErr != nil{
		log.Fatalln("error :", mErr)
	}

	// conncetion is ready and client send data to server
	_, wErr := connection.Write(data)
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