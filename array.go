package main

import (
	"fmt"
)

func main() {
var Nama =  [...]string {
	"Rizfan",
	"Anjay Mabar",
	"Slebeew!",
	"hihihaha",
}
// Nama[3] = "Anjay Mabar" 

fmt.Println(Nama)
fmt.Println()
// fmt.Println(Nama[3])

var values = [...]int{123, 90, 92, 123, 1232, 12398, 292}
values[1] = 12
fmt.Println(values[0])
fmt.Println(values[1])
fmt.Println(len(values))
}