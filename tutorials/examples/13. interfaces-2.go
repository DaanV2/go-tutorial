package main

type (
	Dog   struct{}
	Cat   struct{}
	Mouse struct{}
)

func (d *Dog) Speak() string {
	return "Woof!"
}

func (c *Cat) Speak() string {
	return "Meow!"
}

func (m *Mouse) Speak() string {
	return "Squeak!"
}

Output(&Dog{}) // Woof!
Output(&Cat{}) // Meow!
Output(&Mouse{}) // Squeak!