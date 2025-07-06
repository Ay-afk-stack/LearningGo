package main

import (
	"fmt"
	"time"
)

func main() {
	dayBorn := time.Sunday

	switch {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born on some other days.")
	}
}