package main

import "fmt"

func main() {

	var totalTax float64

	totalTax += salesTax(0.99, 0.075)
	totalTax += salesTax(2.75, 0.015)
	totalTax += salesTax(0.87, 0.02)

	fmt.Println("Sales Tax Total:",totalTax)
}

func salesTax(cost float64, rate float64) float64 {
	return cost * rate
}