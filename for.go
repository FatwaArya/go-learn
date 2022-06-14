package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	slice := []string{"Fatwa", "Arya", "Satya", "Akbar", "uwow"}

	for index, value := range slice {
		fmt.Println(index, value)
	}

	person := make(map[string]string)
	person["name"] = "Fatwa"
	person["address"] = "Blitar"

	for key, value := range person {
		fmt.Println(key, value)
	}
}
