package main

import "fmt"


func Ups() any {
	return "dr.muthu"
}

func main() {
	var kosong = Ups()
	fmt.Println(kosong)
}