package main

import "fmt"

type Animal interface {
	Speak() string
}

func Output(a Animal) {
	fmt.Println(a.Speak())
}