package main

import "fmt"

type Filter func(string) string; // type declaration

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
   fmt.Println("hello!", filteredName)	
}

func filtered(name string) string {
	if name == "anjing" {
		return "..."
	} else if name == "bangsat" {
		return "..."
	} else {
		return name
	}
}


func main() {
	sayHelloWithFilter("rzfan", filtered)
	sayHelloWithFilter("anjing", filtered)
	sayHelloWithFilter("bangsat", filtered)
}