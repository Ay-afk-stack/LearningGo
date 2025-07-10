package main

import "fmt"

func getWeeks () []string {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	days = append(days[6:],days[:6]... )
	return days
}

func main() {
	weekDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	weekDays = weekDays[:len(weekDays) - 1]
	updatedWeekDays := append([]string {"Sunday"}, weekDays...)
	fmt.Println(updatedWeekDays)
	fmt.Println(getWeeks())
}