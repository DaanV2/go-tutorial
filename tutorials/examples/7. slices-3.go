
items = append(items, 6)
items = append(items, 6, 7, 8)
items = append(items, moreItems...)

items = nil
items = append(items, 1, 2, 3)
fmt.Println(items)
