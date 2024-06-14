package main

// Client is a struct that represents a client
type Client struct {
	Name  string // Name is a field of the Client struct
	Age   int    // Age is a field of the Client struct
	admin bool   // admin is a field of the Client struct
}

func foo(c *Client) {
	c.Name = "foo"
	c.Age = 42
	c.admin = true
}