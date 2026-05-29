package main

import "fmt"

// import "fmt"

func MapBaru(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string] string{
			"name": name,
		}
	}
}

func main() {
	Maps := MapBaru("Rizqy Fajrul Syabani") 

	if Maps == nil {
		fmt.Println("data map masih kosong!")
	} else {
		fmt.Println(Maps["name"])
	}

	// fmt.Println(Maps["name"])
}