package main

import "fmt"

func main() {
	numbers := []int16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Printf("%v even\n", number)
		} else {
			fmt.Printf("%v odd\n", number)
		}
	}
}
