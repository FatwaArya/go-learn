package main

import (
	"fmt"
)

func main() {
	name := "Fatwa"

	switch name {
	case "Fatwa":
		fmt.Println("Hello Fatwa")
	case "Fira":
		fmt.Println("Hello Fira")
	default:
		fmt.Println("Boleh kenalan?")

	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Terlalu Panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Terlalu Panjang")
	case length < 5:
		fmt.Println("nama lumayan benar")
	default:
		fmt.Println("nama sudah benar")

	}

}
