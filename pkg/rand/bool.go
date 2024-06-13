package rand

// RandomBool generates a random boolean value
func RandomBool() bool {
	return Float64(1) >= 0.5
}
