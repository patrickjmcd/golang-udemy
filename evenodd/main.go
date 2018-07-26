package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, sl := range slice {
		eo := "odd"
		if sl%2 == 0 {
			eo = "even"
		}
		fmt.Printf("%v is %v\n", sl, eo)
	}
}
