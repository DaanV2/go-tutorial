package main

type CustomerID = string
type CustomerAge = int

func (id CustomerID) Validate() bool {
	return true
}

func (age CustomerAge) Validate() bool {
	return true
}