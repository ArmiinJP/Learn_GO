package main

import (
	"fmt"
	"log"
	"net"
)

func server(){

	// server listening the port to recive request from user
	listener, lErr := net.Listen("tcp", "localhost:2022")
	if lErr != nil{
		log.Fatalln("can't register ip")
	}
	
	var count int
	for {		
		if count > 5 {
			return
		}
		count++

		// server stop to recive conncetion from client
		conn, err := listener.Accept()
		
		defer listener.Close()
		if err != nil{
			log.Println("err in connection request")
			
			continue
		}

		// server Read request data from user
		var request = make([]byte, 1024)
		numberOfByte, rErr := conn.Read(request)
		if rErr != nil{
			log.Println("error when Reading request")

			continue
		}

		// server send message susccefully recive to client
		if _, wErr := conn.Write([]byte("your message is succesfully recived")); wErr != nil{
			log.Println("error when writing response")

			continue
		}
		
		// Print all event from this connection in terminal
		fmt.Printf("client: %s , send request: %s, and size is: %d\n", 
			conn.RemoteAddr(), string(request), numberOfByte)
	}
}

func main(){

	server()
	// defer => close listening server
	for {
		var count int
		count++
		count--
	}
}