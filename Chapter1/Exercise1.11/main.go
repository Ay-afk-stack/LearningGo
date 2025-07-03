package main

import "fmt"

func main() {
	visits := 35
	fmt.Println("First Visit :", visits == 1)

	fmt.Println("Return Visit :", visits != 1)

	fmt.Println("Silver member :", visits >= 10 && visits <= 20)

	fmt.Println("Gold member :", visits >= 21 && visits <= 30)

	fmt.Println("Platinum Member :",visits > 30)
}