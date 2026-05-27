package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) sayHello() {
	man.Name = "Hi, Mr." + man.Name
}

func main() {
	rzfan := Man{"Rizfan"}
	rzfan.sayHello()

	fmt.Println(rzfan.Name)
}