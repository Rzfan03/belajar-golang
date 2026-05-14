package main

import "fmt"

func main() {
	type Umur int
	var umur Umur
	
	fmt.Println("Masukan Umur mu : ")
	fmt.Scan(&umur)

	if !(umur < 18) {
		fmt.Println("dewasa")
	} else {
		fmt.Println("belum dewasa")
	}
}
