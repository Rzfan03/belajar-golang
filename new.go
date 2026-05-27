package main

import "fmt"

type makanan struct {
	name string
}

func main() {
	var makanan1 *makanan = new(makanan)
	var makanan2 *makanan = makanan1


	makanan2.name = "mie jebeew!"

	fmt.Println(makanan1)
	fmt.Println(makanan2)
}