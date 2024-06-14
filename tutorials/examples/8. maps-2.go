value, present := m["k2"]
value, ok := m["k4"]

if v, ok := m["k1"]; !ok {
	m["k1"] = 1
	return 1
} else {
	return v
}