package main

import (
	"fmt"
)

const text = "hello world"
var text2 string = text + "!"

func main() {
    fmt.Println(text, text2)
}