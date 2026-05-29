package hostname

import (
	"fmt"
)
func init() {
	fmt.Println("This is ur Init!")
}

func GetHostName() string {
	return "linux"
}