package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid","employee","address","hours worked","hourly rate","manager"}
	hdr2 := []string{"employee","empid","hours worked","address","manager","hourly rate"}
	result := csvHdrCol(hdr)
	fmt.Println("Result:")
	fmt.Println(result)
	result2 := csvHdrCol(hdr2)
	fmt.Println("Result2:")
	fmt.Println(result2)

}

func csvHdrCol(header []string) map[int]string{
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v){
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	return csvHeadersToColumnIndex
}