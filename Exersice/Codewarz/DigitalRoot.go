package main

import "fmt"

func DigitalRoot(n int) int {
	var remider, sum int
	if n < 10 {
		return n
	} else {
		for n > 0 {
			remider = n % 10
			sum += remider
			n = n / 10
		}
		return DigitalRoot(sum)
	}
}

func main() {
	fmt.Println(DigitalRoot(732885053789454743))
}
