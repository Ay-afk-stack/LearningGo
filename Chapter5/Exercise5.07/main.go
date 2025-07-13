package main

import "fmt"

func main() {
	decrement := decrementer()
	fmt.Println(decrement())
	fmt.Println(decrement())
}

func decrementer() func() int {
	i := 5
	return func() int{
		i-- 
		return i
	}
}