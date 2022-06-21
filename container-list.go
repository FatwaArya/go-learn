package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Fatwa")
	data.PushBack("Aryasatya")
	data.PushBack("Akbar")

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value, " ")
	}
}
