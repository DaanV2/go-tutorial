package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := make(chan string, 10)

	for i := range 1_000_000 {
		go conv(i, data)
	}

	for v := range data {
		fmt.Println(v)
	}
}

func conv(i int, out chan<- string) {
	out <- strconv.Itoa(i)
}