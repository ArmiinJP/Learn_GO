package testSubdir

import "fmt"

var A = 3
func init(){
	fmt.Println("salam in test3")
}
func Te()string {
	test()
	return fmt.Sprint("this func in the test3.go in the testSubdir package")
}
