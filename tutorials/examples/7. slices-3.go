
items = append(items, "derek")
items = append(items, "banana", "apple", "tomato")
items = append(items, moreItems...)

items = nil
items = append(items, "dave", "debbie", "gandalf")
fmt.Println(items)
