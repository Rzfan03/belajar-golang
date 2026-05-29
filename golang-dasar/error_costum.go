package main

import "fmt"

type Error struct {
	Message string
}

func (v *Error) Error() string {
	return v.Message
}

type NotFound struct {
	Message string
}

func (n *NotFound) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &Error{Message: "Error!"}
	}

	if id != "rizfan" {
		return &NotFound{Message: "Not Found!"}
	}
	return nil
}

func main() {
	result := SaveData("rizfan", nil)
	fmt.Println(result)
	
}