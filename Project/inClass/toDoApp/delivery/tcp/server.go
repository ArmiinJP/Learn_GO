package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"todolistapp/delivery/param"
	"todolistapp/repository/memorystore"
	"todolistapp/service/task"
)

func main(){

	listener, lErr := net.Listen("tcp", "127.0.0.1:2023")
	if lErr != nil{
		fmt.Println("error occured:", lErr)
		//sdfasf
		return
	}
	defer listener.Close()
	log.Println("tcp server listning on 127.0.0.1:2023")

	memRepo := memorystore.NewTask()
	taskService := task.NewService(memRepo)

	for {
		conn, aErr := listener.Accept()
		if aErr != nil{
			fmt.Println("error occured:", aErr)

			continue
		}

		var data = make([]byte, 1024)
		
		numberOfByte, rErr := conn.Read(data)
		if rErr != nil{
			fmt.Println("error occured:", rErr)

			continue			
		}

		request := param.Request{}
		jErr := json.Unmarshal(data[:numberOfByte], &request)
		if jErr != nil{
			fmt.Println("error occured:", jErr)

			continue			
		}

		switch request.Command {
		case "create-task":
			response , cErr := taskService.CreatedTask(task.CreatedRequest{
				Title: request.Values.Title,
				DueDate: request.Values.DueDate,
				CategoryID: request.Values.CategoryID,
				AutheticatedUserID: 0,
			})
			if cErr != nil{
				_, wErr := conn.Write([]byte(fmt.Sprintf("add data %v error %s", response, cErr.Error())))
				if wErr != nil{
					log.Println("error occured:", wErr)

					continue						
				}				
			}
			
			fmt.Printf("data Read: %+v, from %s\n", request, conn.RemoteAddr())

			res, mErr := json.Marshal(&response)
			if mErr != nil{
				_, wErr := conn.Write([]byte(mErr.Error()))
				if wErr != nil{
					log.Println("error occured:", wErr)

					continue					
				}

				continue
			} 
			_, wErr := conn.Write(res)
			if wErr != nil{
				log.Println("error occured:", wErr)

				continue				
			}

		case "list-task":
			response, lErr:= taskService.ListTask(task.ListRequest{UserID: 0})
			if lErr != nil{
				fmt.Println("error occured:", lErr)
			}
			fmt.Println(response.Tasks)
		}


		
		
	}
}