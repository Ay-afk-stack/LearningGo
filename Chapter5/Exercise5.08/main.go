package main

import "fmt"

type calc func(x, y int) int

func salary(x int, y int, f calc) int {
	return f(x, y)
}

func developerSalary(hourlyRate, hoursWorked int) int{
	return hourlyRate * hoursWorked
}

func managerSalary(baseSalary int, bonus int) int {
	return baseSalary + bonus
}

func main() {
	developer := salary(50, 2080, developerSalary)
	manager := salary(150000, 25000, managerSalary)

	fmt.Printf("Boss Salary: %d\n", manager)
	fmt.Printf("Developer Salary: %d\n", developer)
}