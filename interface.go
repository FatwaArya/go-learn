package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

type Person struct {
	Name string
}

//jika nama method sama seperti nama method interface, maka kita bisa menggunakan interface untuk mengimplementasikan method
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var Fatwa Person
	Fatwa.Name = "Fatwa"

	SayHello(Fatwa)

	lion := Animal{
		Name: "Arslan",
	}
	SayHello(lion)
}
