package main

import "fmt"

func main() {

var ( 
	// nama        = "Rzfan03"
	TinggiBadan = 175.5
	BeratBadan  = 70.0
)

type Nama string


	
	// fmt.Println(nama)
	fmt.Println(Nama("Rizqy Fajrul Syabani"))
	cetakUmur(25)
	fmt.Println("TinggiBadan:", TinggiBadan)
	fmt.Println("BeratBadan:", BeratBadan)
}
func cetakUmur(umur rune) {
	fmt.Println("Umur:", umur)
}
