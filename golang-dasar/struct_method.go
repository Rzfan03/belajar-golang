package main

import "fmt"

type name struct {
	Name string
}

func (n name) Hello(age int) {
	fmt.Printf("My name is %s and my age is %d", n.Name, age)
}

func main() {
	rizfan := name{Name: "Rizfan"}
	rizfan.Hello(17)
}