package main

import "fmt"

var status bool

func main() {
	status = true

	if status = false; status {
		fmt.Println("Is True!")
	} else {
		fmt.Println("Is False!")
	}

	switch num := 3; num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(1)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("Greater than 3")
	}

}
