package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
//----------------------------------------------
/*
Notice that the s byte slice will be used in every line that involves writing presented in this Go program. 
Additionally, the fmt.Fprintf() function used here can help you write data to your own log files using the format you want.
In this case, fmt.Fprintf() writes your data to the file identified by f1
*/
	s := []byte("Data to write\n")

	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

//----------------------------------------------
/*
In this case, f2.WriteString() is used for writing your data to a file.
*/
	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
//----------------------------------------------
/*
In this case, bufio.NewWriter() opens a file for writing and bufio.WriteString() writes the data.
*/
	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()
//----------------------------------------------
/*
This method needs just a single function call named ioutil.WriteFile() for writing
your data, and it does not require the use of os.Create().
*/
	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
//----------------------------------------------
/*
The last technique uses io.WriteString() to write the desired data to a file.
*/
	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}