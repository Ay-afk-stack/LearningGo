package main

import "fmt"

func main() {
	num := 5
	squareNumber := func(num int) int{
		return num * num
	}
	fmt.Println(squareNumber(num))
}