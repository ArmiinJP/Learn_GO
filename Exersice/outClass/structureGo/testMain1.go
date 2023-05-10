package main

import a "structureGo/testSubdir"
import "fmt"
func inMain(){
	fmt.Println("you are in main")
}
func testtt(){
	fmt.Println(a.B)
	fmt.Println(a.A)

	fmt.Println(a.Te())

	//pp()
}