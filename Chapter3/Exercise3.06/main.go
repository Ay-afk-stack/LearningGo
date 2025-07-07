package main

import "fmt"

func main() {
	logLevel := "デバッグ"
	
	for index, value := range logLevel {
		fmt.Println(index, string(value))
	}
}