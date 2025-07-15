package main

import "fmt"

func sum(nums ...int) int{
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	res := sum(1,2,3,4,5,6,7,8,9,10)
	res2 := sum(numbers...)
	fmt.Println(res)
	fmt.Println(res2)
}