package main

import (
	"errors"
)

func main() {
	input := []string {"abc", "123"}
	var err error

	for _, i := range input {
		_, cerr := getValue(i)
		err = errors.Join(err, cerr)
	}

	var specificError SpecificError
	if errors.Is(err, OutOfRangeError) {
		return
	} else if errors.As(err, &specificError) {
		return
	} else if err != nil {
		panic(err)
	}
}