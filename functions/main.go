package main

import "fmt"

// func add() (string, string, bool){
// return "Golang", "JS", true
// }

func processIt() func(a int) int {
	return func(a int) int{
		return 2
	}
}
func main() {
	fn := processIt()
	res := fn(6)
	fmt.Println(res)
}

