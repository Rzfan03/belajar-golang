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

	// book := make(map[string]string)
	// book["title"] = "Belajar Golang!"
	// book["author"] = "Rzfan03"


	// fmt.Println(book)
	// fmt.Println(book)

	var book = map[string] int {
		"si kancil": 10,
		"si marzuki": 20,
		"si sugiono": 30,
	}

	delete(book, "si kancil")
	fmt.Print(book)
}