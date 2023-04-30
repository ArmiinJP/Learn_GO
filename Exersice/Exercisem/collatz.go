package main

import "fmt"

func CollatzConjecture(n int) (int, error) {

	if n <= 0 {
		return 0, fmt.Errorf("false")
	
	} else if n == 1{
		return 0, nil
	} else {
		if n % 2 == 0{
			n = n / 2
		} else {
			n = 3*n + 1
		}

		if res , err := CollatzConjecture(n); err != nil{
			return -1, fmt.Errorf("%w" ,err)

		} else {
			return res+1 , nil
			
		}
	}
}

func main(){
	fmt.Println(CollatzConjecture(2))
}