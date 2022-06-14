package main

import (
	"fmt"
)

func main() {
	var name = "Fatwa"

	if name == "Fatwa" {
		fmt.Println("Hello Fatwa")
	} else if name == "Fira" {
		fmt.Println("Hello Fira")
	} else {
		fmt.Println("Boleh kenalan?")

	}

	//short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
