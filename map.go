package main

import "fmt"

func main() {
	// person := map[string]string{
	// 	"nama":  "Rzfan03",
	// 	"Alamat": "Nijang",
	// }


	// fmt.Println(person["nama"])
	// fmt.Println(person["Alamat"])
	// fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang!"
	book["author"] = "Rzfan03"


	fmt.Println(book)


	delete(book, "author")
	fmt.Println(book)
}