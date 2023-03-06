//source: https://gosamples.dev/read-file/

//*** Reading entire file

/*return slice of byte -- all file into the memory
The simplest way of reading a text or binary file in Go is to use the ReadFile() function from the os package.
This function reads the entire content of the file into a byte slice,
so you should be careful when trying to read a large file but For small files, this function is best.
for large file, you should read the file line by line or read file in chunks.
*/

package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    content, err := os.ReadFile("README.md")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(content, string(content))
}
//---------------------------------------------------
//*** Read a file in chunks

/* When you have a very large file or don’t want to store the entire file in memory,
 you can read the file in fixed-size chunks. 
 you need to create a byte slice of the specified size (chunkSize in the example)
 as a buffer for storing the subsequent read bytes.

Using Read() method of the File type, we can load the next chunk of the file data. 
The reading loop finishes when an io.EOF error occurs, indicating the end of the file.
*/
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

const chunkSize = 10

func main() {
    // open file
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    //  (remember to close the file after the operation is done, for example, by using defer statement)
    defer f.Close()

    buf := make([]byte, chunkSize)

    for {
        n, err := f.Read(buf)
        if err != nil && err != io.EOF {
            log.Fatal(err)
        }

        if err == io.EOF {
            break
        }

        fmt.Println(string(buf[:n]))
    }
}
//---------------------------------------------------
//Read a file line by line

/*To read a file line by line, we can use a convenient bufio.Scanner structure.
Its constructor, NewScanner(),
takes an opened file and lets you read subsequent lines through Scan() and Text() methods.
Using Err() method, you can check errors encountered during file reading.
*/

package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    // open file
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        // do something with a line
        fmt.Printf("line: %s\n", scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
//---------------------------------------------------
//Read a file word by word

/*Reading a file word by word is almost the same as reading line by line.
You only need to change the split function of the Scanner from the default ScanLines() to ScanWords().
*/
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    // open file
    f, err := os.Open("file.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file word by word using scanner
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        // do something with a word
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}