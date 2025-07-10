package main

import (
	"fmt"
	"os"
)

var users = map[string]string {
	"305" : "Sue",
	"204" : "Bob", 
	"631" : "Jake",
	"073" : "Tracy",
}

func deleteUser(id string) {
	delete(users, id)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID must be greater than 2")
		os.Exit(1)
	}
	id := os.Args[1]
	deleteUser(id)
	fmt.Println("Users:",users)

}