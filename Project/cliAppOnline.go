package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)
 
func CreateFile() {
    fmt.Printf("Writing to a file in Go lang\n")
    file, err := os.Create("test.txt")
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
    defer file.Close()
    len, err := file.WriteString("test,123,flaf,gg,read\n")
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    len, err = file.WriteString("test,123,flaf,gg,read\n")
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    len, err = file.WriteString("test,123,flaf,gg,read\n")
    if err != nil {
        log.Fatalf("failed writing to file: %s", err)
    }
    fmt.Printf("\nFile Name: %s", file.Name())
    fmt.Printf("\nLength: %d bytes", len)
}
func ReadFile() {
 
    fmt.Printf("\n\nReading a file in Go lang\n")
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nSize: %d bytes", len(data))
    fmt.Printf("\nData: %s", data)
 
}
func Readfile2(){
	readFile, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
 
	for fileScanner.Scan() {
		fmt.Printf("%s",strings.Split(fileScanner.Text(), ",")[1])
		break	
	}
 
	readFile.Close()
}
func main(){
    //CreateFile()
    //Readfile2()
	var flagRegion string
	fmt.Scanln(&flagRegion)
	fmt.Println(flagRegion)
}