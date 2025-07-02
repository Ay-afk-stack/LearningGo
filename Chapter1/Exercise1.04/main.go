package main

import (
	"fmt"
	"time"
)
var (
	Debug    bool  
	LogLevel string = "info"
	startUpTime time.Time = time.Now()
)
func main() {
		var hello int
 		fmt.Println(hello, Debug, LogLevel, startUpTime)
}