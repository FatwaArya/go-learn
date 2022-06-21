package main

import (
	"fmt"
	"go-learn/helper"
	"strconv"
)

func main() {
	boolen, err := strconv.ParseBool("salah")
	helper.Check(boolen, err)

	number, err := strconv.ParseInt("100000", 10, 64)
	helper.Check(number, err)

	value := strconv.FormatInt(100000, 10)
	fmt.Println(value)
}
