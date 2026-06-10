package main

import "fmt"

type Spp interface {
	CheckSpp() string
}

func CheckSiswa(s Spp) {
	fmt.Println(s.CheckSpp())
}

type Siswa struct {
	Nama string
}

func (s Siswa) CheckSpp() string {
	return "Okelah " + s.Nama + " dah bayar kau yaa!"
}

func main() {
	siswa := Siswa{Nama: "saleh"}
	CheckSiswa(siswa)
}
