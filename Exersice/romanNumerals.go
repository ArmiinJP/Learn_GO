package main

import (
	"fmt"
	"errors"
)

var Units = map[byte]string{'1': "I", '2': "II", '3': "III", '4': "IV", '5': "V", '6': "VI", '7': "VII", '8': "VIII", '9': "IX", '0': ""}
var Tens = map[byte]string{'1': "X", '2': "XX", '3': "XXX", '4': "XL", '5': "L", '6': "LX", '7': "LXX", '8': "LXXX", '9': "XC", '0': ""}
var Hundreds = map[byte]string{'1': "C", '2': "CC", '3': "CCC", '4': "CD", '5': "D", '6': "DC", '7': "DCC", '8': "DCCC", '9': "CM", '0': ""}	
var Thousands = map[byte]string{'1': "M", '2': "MM", '3': "MMM", '0': ""}

func ToRomanNumeral(input int) (string, error) {

	if input <= 0 || input >= 4000 {
		return "", errors.New("input is False")
	}

	result := ""
	inputString := fmt.Sprint(input)

	for i := len(inputString)-1 ; i > -1; i--{
		switch i {
		case 3: result += Thousands[inputString[len(inputString)-1-i]]
		case 2: result += Hundreds[inputString[len(inputString)-1-i]]
		case 1: result += Tens[inputString[len(inputString)-1-i]]
		case 0:	result += Units[inputString[len(inputString)-1-i]]			
		default: return "", errors.New("input is False")
		}
	}
	return result, nil
}


func main(){
	fmt.Println(ToRomanNumeral(1990))
}