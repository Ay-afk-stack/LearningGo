package main

import (
	"errors"
	"fmt"
)

func typeCheck(data any) (string,error){
		switch t := data.(type) {
		case int, int32, int64 :
			if v, ok := t.(int); ok {
				return fmt.Sprint(v," is int"),nil
			}
			if v, ok := t.(int); ok {
				return fmt.Sprint(v," is int"),nil
			}
			if v, ok := t.(int); ok {
				return fmt.Sprint(v," is int"),nil
			}
		case float32, float64:
			if f, ok := t.(float64); ok {
				return fmt.Sprint(f," is float"), nil
			}
			return fmt.Sprint(t.(float32), " is float"),nil
		case string:
			return t + " is string", nil
		case bool:
			if t{
				return "true is bool",nil
			}
			return "false is bool",nil
		default:
			return "Unknown type", nil
	}
	return "", errors.New("Unsupported Type assertion")
}

func main() {
	data := []any{1, 3.14, "hello", true, struct{}{}}
	for _, val := range data{
		res,_ := typeCheck(val)
		fmt.Println(res)
	}
}