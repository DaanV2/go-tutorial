package main

func divide(a, b int) (int, int) {
	div := a / b
	mod := a % b
	return div, mod
}
func divide(a, b int) (div int, mod int) {
	div = a / b
	mod = a % b
	return
}

a, b := divide(10, 3)
_, c := divide(10, 3)
d, _ := divide(10, 3)
_, _ := divide(10, 3)
