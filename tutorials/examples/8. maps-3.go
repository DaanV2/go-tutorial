n := map[string]int{"foo": 1, "bar": 2}
n2 := map[string]int{"foo": 1, "bar": 2}
if maps.Equal(n, n2) {
	fmt.Println("n == n2")
}

for key, value := range n {
	fmt.Println(key, value)
}