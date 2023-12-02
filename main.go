package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	contents := string(file)
	result := trebuchet(contents)
	fmt.Println(result)
}
