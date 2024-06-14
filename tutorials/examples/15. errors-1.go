package main

import (
	"fmt"
	"strconv"
)

func main() {
	if v, err := getValue("abc"); err != nil {
		fmt.Println("error during conversion:", err)
		return
	} else {
		fmt.Println("value:", v)
	}	
}

func getValue(input string) (int, error) {
	v, err := strconv.ParseInt(input, 10, 64)
	return int(v), err
}