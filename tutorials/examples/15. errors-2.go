package main

import (
	"errors"
	"fmt"
	"strconv"
)

type SpecificError struct {
	Msg string
	Value string
}

func (e SpecificError) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.Value)
}

var OutOfRangeError = SpecificError{"value out of range", ""}

func main() {
	v, err := getValue("abc")
	var specificError SpecificError
	if errors.Is(err, OutOfRangeError) {
		return
	} else if errors.As(err, &specificError) {
		return
	} else if err != nil {
		fmt.Println("error during conversion:", err)
		return
	}
}

func getValue(input string) (int, error) {
	v, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, SpecificError{"error during conversion", input}
	}

	return int(v), err
}