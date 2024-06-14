m := make(map[string]int, 100)

m["k1"] = 7
m["k2"] = 13

v1 := m["k1"]
v3 := m["k3"]

amount := len(m)

delete(m, "k2")
clear(m)