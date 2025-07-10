package main

import (
	"fmt"
	"os"
)

type locale struct {
	language string
	country string
} 

func getLocales() []locale {
	locales := []locale{
		{"en_US","USA"},{"en_CN","Canada"},{"fr_CN","Canada"},{"fr_FR","France"},
		{"ru_RU","Russia"},
	}
	return locales
}

func main() {
	locales := getLocales()

	if len(os.Args) < 2 {
		fmt.Println("Argument must be greater than 2.")
		os.Exit(1)
	}

	localeLanguage := os.Args[1]
	for _, val := range locales {
		if localeLanguage == val.language{
			fmt.Println("Locale is supported")
			break
		} else {
			fmt.Printf("Locale is not supported: %v\n", localeLanguage)
			os.Exit(1)
		}
	}
	
}