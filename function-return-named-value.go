package main

import "fmt"

func getFullName2() (firstName string, middleName string) {
	firstName = "Fatwa"
	middleName = "Arya"
	return

}

func main() {
	firstName, _ := getFullName2()
	fmt.Println("Hello", firstName)
}
