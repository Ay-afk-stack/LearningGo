package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error){
	switch t := v.(type){
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		}
		return "falsefalse", nil
	case float32, float64:
		if f, ok := t.(float64); ok{
			return fmt.Sprintln(f * 2),nil
		}
		return fmt.Sprintln(t.(float32) * 2),nil
	case int,int32, int64:
		if i, ok := t.(int64); ok {
			return fmt.Sprintln(i * 2), nil
		} else if i, ok := t.(int); ok {
			return fmt.Sprintln(i * 2),nil
		}
		return fmt.Sprintln(t.(int32) * 2),nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	res, _ := doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("Yum:", res)
	res, _ = doubler(false)
	fmt.Println("True:",res)
	res, _ = doubler(3.14)
	fmt.Println("True:",res)
}