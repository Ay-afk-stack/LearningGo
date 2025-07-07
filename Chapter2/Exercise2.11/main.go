package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for {
		r := rand.Intn(8)
		if r % 3 == 0 {
			fmt.Println("SKIP!!!")
			continue
		} else if r % 2 == 0 {
			fmt.Println("STOP!!!")
			break
		} else {
			fmt.Println(r)
		}
	}
}