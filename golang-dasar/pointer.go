package main

import "fmt"

type makanan struct {
	name string
}

func main() {
	makanan1 := makanan{"Mie yamin"}
	makanan2 := makanan1

	makanan2.name = "Mie goyang"

	fmt.Println(makanan1)
	fmt.Println(makanan2)
 nama := "Rizfan03"
 nama2 := &nama
 *nama2 = "rzfan03"

 fmt.Println(&nama)
 fmt.Println(nama2)
 fmt.Println(nama)
 }