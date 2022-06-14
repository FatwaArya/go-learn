package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "ini kata kasar"
	}
	return "Hello " + name
}

func main() {
	sayHelloWithFilter("Fatwa", spamFilter)

	sayHelloWithFilter("Anjing", spamFilter)

}
