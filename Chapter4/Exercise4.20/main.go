package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error){
	if i, ok := v.(int); ok {
		return fmt.Sprintln(i * 2), nil
	}
	if s, ok := v.(string); ok {
		return s + s, nil
	}
	return "", errors.New("unsupported type passed!")
}

func main() {
	res, _ := doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("Yum:", res)
	_, err := doubler(true)
	fmt.Println("True:",err)
}