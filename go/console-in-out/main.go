package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num int
	var num2 int
	var text string

	fmt.Print("Ingrese número 1: ")
	fmt.Scanf("%d", &num)

	fmt.Print("Ingrese número 2: ")
	fmt.Scanf("%d", &num2)

	fmt.Printf("%d", num+num2)

	fmt.Print("Texto: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text = scanner.Text()
	}

	result := num + num2
	fmt.Println(text, result)

}
