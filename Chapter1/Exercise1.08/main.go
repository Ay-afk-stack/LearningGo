package main

import "fmt"

func main() {
	query, limit, offset := "bat", 10, 0

	fmt.Println(query, limit, offset)

	query, limit, offset = "ball", offset, 10

	fmt.Println(query, limit, offset)

	query = "haha"
	fmt.Println(query)

}