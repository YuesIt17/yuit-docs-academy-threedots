package main

import (
	"fmt"
	"os"
)

func main() {
	if _,err := os.Stat("analysis.xlsx"); err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("ok")
	}
}
