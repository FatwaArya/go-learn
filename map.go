package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Fatwa",
		"address": "Blitar",
	}
	person["title"] = "Petani"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book = make(map[string]string)
	book["title"] = "Molecule of more"
	book["author"] = "Fatwa"
	book["wrong"] = "ups"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))
}
