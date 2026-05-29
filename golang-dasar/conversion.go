package main

import "fmt"

func main() {
	var Umur32 int32 = 32768
	var Umur64 int64 = int64(Umur32)
	var Umur16 int16 = int16(Umur64)

	fmt.Println(Umur32, Umur64, Umur16)
	
}
