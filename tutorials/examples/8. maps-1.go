m := make(map[string]string, 100)

m["k1"] = "steve"
m["k2"] = "jane"

v1 := m["k1"]
v3 := m["k3"]

amount := len(m)

delete(m, "k2")
clear(m)