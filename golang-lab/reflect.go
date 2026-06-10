package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 200
	var resultNumber = reflect.ValueOf(number)

	fmt.Println(resultNumber.Type())
}