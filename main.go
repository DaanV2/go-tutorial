/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(chan uint64)

	wg := sync.WaitGroup{}

	for i := range 1_000_000 {
		wg.Add(1)

		go func(i uint64, data chan <- uint64) {
			defer wg.Done()
			data <- i * 2
		}(uint64(i), data)
	}

	go func() {
		fmt.Println("Im awaiting the group")
		defer fmt.Println("Group has been awaited")
		wg.Wait()
		close(data)
	}()

	fmt.Println("Data is ready to be consumed")
	defer fmt.Println("Data has been consumed")

	sum := uint64(0)
	for i := range data {
		sum += i
	}

	fmt.Println(sum)
}