package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	database := flag.String("user", "root", "put your database user")
	password := flag.String("password", "root", "put your database password")

	flag.Parse()

	fmt.Println(*host)
	fmt.Println(*database)
	fmt.Println(*password)
}
