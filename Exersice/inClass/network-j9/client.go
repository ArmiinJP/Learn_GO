package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main(){

	// recive message from user to send to server
	message := flag.String("message", "empty", `message you want to send server`)
	flag.Parse()
	
	//client start connect to server
	connection, cErr := net.Dial("tcp", "127.0.0.1:2022")
	if cErr != nil{
		log.Fatalln("error connection")
	}
	
	// conncetion is ready and client send data to server
	numberOfByte, wErr := connection.Write([]byte(*message))
	if wErr != nil {
		log.Fatalln("error when writeing buffer")
	}
	// client stop to recive reponse from server that data succefully recive from server
	var response = make([]byte, 1024)
	if _, rErr := connection.Read(response); rErr != nil{
		log.Fatalln("error when getting response")
	}
	
	//Print all event in terminal
	fmt.Printf("your address is: %s, number of byte you send is: %d, response is: %s\n",
	 connection.LocalAddr(), numberOfByte, response)

}