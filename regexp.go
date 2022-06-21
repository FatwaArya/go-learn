package main

import (
	"fmt"
	"reflect"
)

func main() {
	sample := Sample{Name: "Fatwa"}
	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))

	sample.Name = ""
	fmt.Println(isValid(sample))
}
