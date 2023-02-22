package main

import (
	"fmt"
	"strings"
)

func ToRNA(dna string) string {

    dna = strings.Replace(dna,"C","*",-1)
	dna = strings.Replace(dna,"G","C",-1)
    dna = strings.Replace(dna,"*","G",-1)
    
	dna = strings.Replace(dna,"A","*",-1)
    dna = strings.Replace(dna,"T","A",-1)
    dna = strings.Replace(dna,"*","U",-1)
    return dna
}


func main(){
	fmt.Println(ToRNA("CATT"))
}