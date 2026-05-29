package main

// import (
// 	"belajar-golang/helper"
// 	"fmt"
// )

// func main() {
// 	result := helper.sayHello("rizfan")
// 	fmt.Println("hi")
// 	fmt.Println(result)
// }

import (
	"belajar-golang/database"
	"belajar-golang/hostname"
	"fmt"
)

func main() {
	hostname := hostname.GetHostName()
	result := database.GetNameDatabase()
	fmt.Println(result, "\n","Your system is : ",hostname)
}