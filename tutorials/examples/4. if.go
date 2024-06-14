package main

func foo(i int) bool {
	if i == 0 {
		return true
	}

	if i == 0 {
		return true
	} else {
		return false
	}

	if i == 0 || i == 1 {
		return true
	}

	if num := i * 25; num < 0 {
		return true
	} else if num < 10 {
		return false
	}
}
