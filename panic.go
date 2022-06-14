package main

import "fmt"

func endApp() {
	fmt.Println("Applikasi selesai")
	message := recover()
	if message != nil {

		fmt.Println("Error dengan message: ", message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("Run App")

}

func main() {
	runApp(false)
}
