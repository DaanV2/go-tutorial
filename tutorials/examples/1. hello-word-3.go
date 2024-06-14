package main

import (
	"fmt"
	"strings"
)

const (
    text = "hello world"
    text2 string = text + "!"
)
var (
    text3 = text
    text4 string = strings.ToUpper(text)
)

func main() {
    fmt.Println(text, text2)
}