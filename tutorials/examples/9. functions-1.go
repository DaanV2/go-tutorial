package main

func plus(a int, b int) int {
	return a + b
}
func plus(a, b, c int) int {
	return a + b + c
}
func plus(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}