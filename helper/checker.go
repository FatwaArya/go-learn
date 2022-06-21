package helper

import "fmt"

func Check(dataType interface{}, err interface{}) {
	if err == nil {
		fmt.Println(dataType)
	} else {
		fmt.Println("Error", err)
	}

}
