package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Broo"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Fatwa")
	fmt.Println(result)
}
