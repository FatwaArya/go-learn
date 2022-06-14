package main

import (
	"fmt"
)

func main() {

	type NoKTP string
	type Married bool
	var noKtpFatwa NoKTP = "159456789159"
	var marriedStatus Married = false

	fmt.Println(noKtpFatwa)
	fmt.Println(marriedStatus)

}
