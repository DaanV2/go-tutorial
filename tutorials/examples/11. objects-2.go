package main

func foo(c *Client) {
	c.Name = "foo"
	c.Age = 42
	c.admin = true
}

func main() {
	c := Client{Name: "bar", Age: 41, admin: false}
	c := Client{"bar", 41, false}
	foo(&c)
}
