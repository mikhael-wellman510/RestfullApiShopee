package helper

import "fmt"

func PanicIfErr(err error) {
	if err != nil {
		fmt.Println("Error Panic")
		panic(err)
	}
}
