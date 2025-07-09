package main

import "fmt"

func fillArray(arr [10]int) [10]int {
	for a := 0 ; a < len(arr) ; a++ {
		arr[a] = a + 1
	}
	return arr
}

func main() {
	var arr [10]int
	arr = fillArray(arr)
	fmt.Println(arr)
}
