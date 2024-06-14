package examples

func WhatAnimal(a Animal) string {
	if d, ok := a.(*Dog); ok {
		return "Dog"
	}
}