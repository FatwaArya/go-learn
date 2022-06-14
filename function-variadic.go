package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sumAll(1, 2, 3, 4, 5)
	fmt.Println(result)

	slice := []int{1, 2, 3, 4, 805}
	result = sumAll(slice...)
	fmt.Println(result)
}
