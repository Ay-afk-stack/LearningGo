package main

import "fmt"

func main() {
	s1 := []string{
		"Good",
		"Good",
		"Bad",
		"Good",
		"Good",
	}
	for i, val := range s1 {
		if val == "Bad"{
			s1 = append(s1[:i],s1[i+1:]...)
		}
	}
	fmt.Println(s1)
}