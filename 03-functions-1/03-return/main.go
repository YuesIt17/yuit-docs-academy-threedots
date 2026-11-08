package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}