package main

import (
	"fmt"
	"reflect"
	"reflectTest/richError"
	"time"
)
func main() {
	var r = richError.RichError{ 
				Message: "test",
				Operation: "main",
				MetaData: map[string]string{
						"user": "test",
						"id": "2342",
				},
				Time: time.Now(),
			}
	
	
	value := reflect.ValueOf(r)
	fmt.Println("kind of value is: ", value.Kind())


	switch value.Kind(){
	case reflect.Struct:
		fmt.Println("type of value is: ", value.Type())
		fmt.Println("number of field is: ", value.NumField())
		fmt.Println("number of method is: ", value.NumMethod())

		for i := 0; i < value.NumField(); i++{
			fmt.Println("--------------------------------------------")
			fmt.Printf("Kind of field %d is: %s\n",i, value.Field(i).Kind().String())
			fmt.Printf("Type of field %d is: %s\n",i, value.Field(i).Type().String())
			fmt.Printf("Name of field %d is: %s\n",i, value.Type().Field(i).Name)
			fmt.Printf("value of field %d is: %s\n",i, value.Field(i))
		}
	}
	
	

}