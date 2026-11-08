package main

import (
	"fmt"
	"os"
)

func CheckFile(filename string) bool {
	if _, err := os.ReadFile(filename); err != nil {
		return false
	}
	return true
}

func main() {
	ok := CheckFile("input.csv")
	if ok {
		fmt.Println("File correctly read")
	} else {
		fmt.Println("Failed to read file")
	}
}
