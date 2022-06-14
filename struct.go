package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi() {
	fmt.Println("Hello ", customer.Name)
}

func main() {
	fatwa := Customer{
		Name:    "Fatwa",
		Address: "Jakarta",
		Age:     18,
	}

	var Fira Customer
	Fira.Name = "Fira"
	Fira.Address = "Jakarta"
	Fira.Age = 18
	Fira.sayHi()
	fatwa.sayHi()
}
