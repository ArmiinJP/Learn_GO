package main

import (
	"log"
	"net"
)

func main(){
	
	listener, lErr := net.Listen("tcp", ":2222")
	if lErr != nil{
		log.Fatalln("listening to the port refused: ", lErr.Error())
	}
	
	for {
		_, aErr := listener.Accept()
		if aErr != nil{
			log.Println("Accept Connection Error: ", aErr.Error())
			
			continue
		}
		

	}

}