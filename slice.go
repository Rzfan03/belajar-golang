package main

import "fmt"


func main() {
	names := [...]string{"Rizfan", "Anjay", "Mabar", "Slebeew", "Pecah", "Muncung!"}
	
	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)


	fmt.Print(cap(slice2))

	fmt.Println("Tambah Data baru?")
}