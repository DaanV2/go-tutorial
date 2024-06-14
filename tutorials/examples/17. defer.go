package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Closing file")
	}()

	f, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
}
