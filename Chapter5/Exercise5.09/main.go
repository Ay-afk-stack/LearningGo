package main

import "fmt"

func main() {
	defer done()
	fmt.Println("Main: started")
	fmt.Println("Main: end")
}

func done() {
	fmt.Println("Main: Done")
}