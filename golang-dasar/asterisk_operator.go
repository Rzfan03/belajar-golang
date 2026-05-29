package main

import "fmt"

type makanan struct {
	name string
}

func main() {
	makanan1 := makanan{"Mie yamin"}
	makanan2 := &makanan1

	makanan2.name = "Mie goyang"
	makanan2 = &makanan{"chicken barbeque"}

	*makanan2 = makanan{"rendang..."}
	// fmt.Println(makanan1)
	fmt.Println(makanan2)
}