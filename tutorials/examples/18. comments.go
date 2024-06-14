// main package has examples shown
package main

// Data is a struct that contains the unique identifier of the data
type Data struct {
	ID  int // ID is the unique identifier of the data
}

// String returns a string representation of the data
func (d *Data) String() string {
	return "Data with ID: " + string(d.ID)
}