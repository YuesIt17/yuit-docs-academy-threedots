package main

import "fmt"

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}

func Swap(x, y string) (string, string) {
	if x == "" || y == "" {
		return "", ""
	}
	return y, x
}	