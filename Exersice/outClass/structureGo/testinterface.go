package main

import "fmt"


type test struct{
	F0 string
	F1 interface{}
}

type field1 struct{
	A0 string
	a1 string
}

type field2 struct{
	a2 string
	a3 string
}

func main(){
	tmp := test{
		F0: "hasan",
		F1: field1{
			A0: "reza",	
			a1: "ali",
		},
	}

	fmt.Println(tmp)
}