package main

import (
	"fmt"
	"strconv"
)

func main() {
	v, err := getValue("abc")
	if err != nil {
		fmt.Println("error during conversion:", err)
		return
	}

	fmt.Println("value:", v)
}

func getValue(input string) (int, error) {
	v, err := strconv.ParseInt(input, 10, 64)
	return int(v), err
}