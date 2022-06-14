package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "is blacklisted")
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	registerUser("Fatwa", func(name string) bool {
		return name == "Anjing"
	})

}
