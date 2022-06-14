package main

import (
	"fmt"
)

func factorialLoop(value int) int {
	result := 1
	for i := 1; i <= value; i++ {
		result *= i
	}
	return result
}
func factorialRecursive(value int) int {
	if value == 0 { //kondisi berhenti
		return 1
	}
	return value * factorialRecursive(value-1)
}

func main() {
	fmt.Println(factorialLoop(5))

}
