package main

import (
	"fmt"
	"time"
)
var (
	Debug    bool   = false
	LogLevel string = "info"
	startUpTime time.Time = time.Now()
)
func main() {
		var hello = 2
		fmt.Println(hello, Debug, LogLevel, startUpTime)
}