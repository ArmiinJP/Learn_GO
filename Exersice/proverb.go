package main

import "fmt"

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT = iota	//not a triangle
    Equ			//equilateral
    Iso			//isosceles
	Sca			//scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if (a > 0 && b > 0 && c > 0) && (a + b > c && b + c > a && a + c > b) {
		if a == b && a == c && b == c {
			k = Equ
		} else if a == b && a == c || b == c && b == a || c == a && c == b {
			k = Iso
		} else {
			k = Sca
		}
	} else {k = NaT}

	return k
}

func main(){
	fmt.Println()
}

// An equilateral triangle has all three sides the same length.
// An isosceles triangle has at least two sides the same length. (It is sometimes specified as having exactly two sides the same length, but for the purposes of this exercise we'll say at least two.)
// A scalene triangle has all sides of different lengths.
