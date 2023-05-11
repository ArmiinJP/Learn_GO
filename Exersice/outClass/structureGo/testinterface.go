package main

import (
	"fmt"
	"time"
)

type test struct {
	F0 string
	F1 interface{}
}

type field1 struct {
	A0 string
	a1 string
}

type field2 struct {
	a2 string
	a3 string
}

func main() {
	tmp := test{
		F0: "hasan",
		F1: field1{
			A0: "reza",
			a1: "ali",
		},
	}

	dateString := "2021-11-22"
	date, error := time.Parse("2006-01-02", dateString)

	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Printf("Type of dateString: %T\n", dateString)
	fmt.Printf("Type of date: %T\n", date)
	fmt.Println()
	fmt.Printf("Value of dateString: %v\n", dateString)
	fmt.Printf("Value of date: %v", date)

	fmt.Println(tmp)
}
