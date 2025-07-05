package main

import "fmt"

const (
	Pending = iota
	_
	Approved
	Rejected
)

func main() {
	fmt.Println("Pending:", Pending)
	fmt.Println("Approved:", Approved)
	fmt.Println("Rejected:", Rejected)
}