package main

import "fmt"

type Student struct {
	ID int
	Name string
	Jurusan string
}

func main() {
	var std Student
	std.ID = 1
	std.Name = "Alex"
	std.Jurusan = "Layanan Pariwisata"

	
	std2 := &std

	std2.ID = 2
	std2.Name = "Rizqy Fajrul Syabani"
	std2.Jurusan = "Rekayasa Perangkat Lunak"

	fmt.Println(std)
	fmt.Println(std2)
}
