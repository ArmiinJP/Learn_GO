package main

import (
    "fmt"
    "strconv"
)

func main() {
    if i, err := strconv.ParseInt("00110100010010111111111111111111111001101111101001", 2, 64); err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%d\n", i)
    }
}