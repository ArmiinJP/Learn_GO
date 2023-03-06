//Appending to file if exist, Create a file in non exist

/*
Golang allows us to append or add data to an already existing file. 
We will use its built-in package called os with a method OpenFile() to append text to a file.
The os.OpenFile() function takes the following parameters:

os.OpenFile(name/path string, flag int, perm FileMode)

name/path: The name of the file, e.g., “example.txt,” if the file is in the current directory or the complete path of the file if it is present in another directory. The data type of this parameter is string.

flag: An int type instruction given to the method to open the file, e.g., read-only, write-only, or read-write. Commonly used flags are as follows:
O_RDONLY: It opens the file read-only.
O_WRONLY: It opens the file write-only.
O_RDWR: It opens the file read-write.
O_APPEND: It appends data to the file when writing.
O_CREATE: It creates a new file if none exists.

perm: A numeric value of the mode that we want os.OpenFile() to execute in, e.g., read-only has a value of 4 and write-only has a value of 2.

return Value:
File: A File on which different operations such as write or append can be performed based on the file mode passed to the function
*PathError: An error while opening or creating the file.
*/

package main

import(
  "fmt"
  "os"
) 

func main() {

    file,err := os.OpenFile("example.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    
    if err != nil {
		  fmt.Println("Could not open example.txt")
      return
	  }

	  defer file.Close()
	 
    _, err2 := file.WriteString("Appending some text to example.txt")

	  if err2 != nil {
		  fmt.Println("Could not write text to example.txt")
      
	  }else{
      fmt.Println("Operation successful! Text has been appended to example.txt")
    }
}

/*
Line 10: We open example.text using os.OpenFile() with the os.O_Append flag because we want to append data to the file.
os.OpenFile() allows us to provide multiple flags for efficiency by using the OR(|) operator.
Here, we provide the os.O_CREATE flag if example.txt does not exist.
Since we write to the file, the os.O_WRONLY flag specifies write-only mode.
0644 is the numerical representation of all these flags.
*/
