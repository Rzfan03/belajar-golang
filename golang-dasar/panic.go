package main

import "fmt"

func RunApp() {
	fmt.Println("Program Running!")
	pesan := recover()
	fmt.Println("Aduhh telah terjadi panic!: ", pesan)
}

func errApp(error bool) {
	defer RunApp()

	if error {
		panic("Ups Error!!!")
	}
}

func main() {
	errApp(true)
}
