package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("Pre:", a, b)
	// Call swap here
	swap(&a, &b)
	fmt.Println("Post:", a, b)
	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	//swap the values here
	*a, *b = *b, *a
}