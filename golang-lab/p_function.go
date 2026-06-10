package main

import "fmt"

type Food struct {
	Name string
	Price uint
}

func OrderFood(f *Food) {
	f.Name = "ayam geprek"
	f.Price = 2500
}


func main() {
	food := Food{}
	OrderFood(&food)
	fmt.Println(food)
}