package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna" : 3,
		"You" : 3,
		"Give" : 2,
		"Never" : 1,
		"Up" : 4,
	}

	var word string
	var count int
	
	for key, value := range words{
		if value > count {
			word = key
			count = value
		}
	}
	fmt.Println("Most Popular Word:",word)
	fmt.Println("With a count of:", count)
}