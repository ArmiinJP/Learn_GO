
package main

import (
	"fmt"
	"strings"
)

func Valid(id string) bool {

	id = strings.Replace(id, " ", "", -1)	
	
	//delete for small length
	if len(id) <= 1 {
		return false
	}

	for _, v := range id{
		if (rune(v) >= 48) && (rune(v) <= 57){
			continue
		} else {
			return false
		}
	}

// function isValid(cardNumber[1..length])
//     sum := 0
//     parity := length mod 2
//     for i from 1 to length do
//         if i mod 2 != parity then
//             sum := sum + cardNumber[i]
//         elseif cardNumber[i] > 4 then
//             sum := sum + 2 * cardNumber[i] - 9
//         else
//             sum := sum + 2 * cardNumber[i]
//         end if
//     end for
//     return sum mod 10 = 0
// end function

	sum := 0
	parity := len(id) % 2
	for i := range id{
		tmp := int(id[i] - '0')
		if i % 2 != parity{
			sum += tmp
		} else if tmp > 4 {
			sum = sum + 2 * tmp - 9
		} else {
			sum = sum + 2 * tmp
		}
	}
	return sum % 10 == 0
}

func main(){
	fmt.Println(Valid(""))
}