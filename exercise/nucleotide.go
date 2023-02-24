package main

import (
	"fmt"
	"errors"
)

type Histogram map[byte]int
type DNA string

func (d DNA) Counts() (Histogram, error) {
    h := Histogram {'A': 0, 'C': 0, 'G': 0, 'T': 0}
    for i := range d{
		if _, ok := h[d[i]]; ok{
			h[d[i]]++
		} else {
			return Histogram{}, errors.New("input is invalid")
		}  
    }
	return h, nil
}

func main(){
	var tmp DNA = "GATGACA"
	fmt.Println(tmp.Counts())
}