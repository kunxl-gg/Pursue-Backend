package helpers

import "fmt"

func CheckError(err error) {
	fmt.Println(err)
	panic(err)
}

func CheckErrorWithCustomMessage(err error) {
	panic(err)
}
