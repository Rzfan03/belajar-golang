package main

import "fmt"

	type Siswa struct {
		name, alamat string
		nis, umur int
	}
func main() {

	var S Siswa

	S.name = "rzfan03"
	S.alamat = "irian lah boz!"
	S.umur = 17
	S.nis = 1234567

	fmt.Println(S)


	joko := Siswa {
		name: "rizqy fajrul syabani",
		alamat: "irian",
		umur: 17,
		nis: 12312312,
	}

	fmt.Println(joko)


	budi := Siswa{"rizqy fajrul syabani", "irian", 17, 131232}
	fmt.Println(budi)

	budi.katakanPeta("rizfan")

	
}
func (S Siswa) katakanPeta(name string) {
	fmt.Println("Hi", name, "my name is", S.name)
}