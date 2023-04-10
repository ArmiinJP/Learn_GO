package main

import "fmt"
import "strings"

func Score(word string) int {
	totalPoint := 0
	word = strings.ToUpper(word)
	points := map[string]int{
			"A, E, I, O, U, L, N, R, S, T": 1,
			"D, G": 2,
			"B, C, M, P": 3,
			"F, H, V, W, Y": 4,
			"K": 5,
			"J, X": 8,
			"Q, Z": 10,
			}
	for _, char := range word{
		for key, value := range points{
			if strings.Contains(key, string(char)){
				totalPoint += value
			}
		}
	}
	return totalPoint
}

func main() {

	fmt.Println(Score("cabbage"))
}
