package main

import "fmt"

func getFullName() (string, string) {
	return "Fatwa", "Arya"
}

func main() {
	firstName, _ := getFullName()
	fmt.Println("Hello", firstName)
}
