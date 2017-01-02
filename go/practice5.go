package main

import (
	"fmt"
	"strconv"
	"os"
)

func calc(a int, b int, operator string) int {
	if operator == "+" {
		return a + b
	} else if operator == "-" {
		return a - b
	} else if operator == "*" {
		return a * b
	} else if operator == "/" {
		return a / b
	} else {
		return 0
	}
}

func main() {
	var first, second string
	fmt.Print("Waht is the first number? ")
	fmt.Scanln(&first)
	fmt.Print("Waht is the second number? ")
	fmt.Scanln(&second)
	
	f, err := strconv.Atoi(first)
	if err != nil {
		fmt.Println("First is not a Integer")
		os.Exit(1)
	}
	s, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println("Second is not a Integer")
		os.Exit(1)
	}
	
	if f < 0 || s < 0 {
		fmt.Println("음수를 넣지 마세여")
		os.Exit(1)
	}
	
	fmt.Printf("%s + %s = %d\n", first, second, calc(f, s, "+"))
	fmt.Printf("%s - %s = %d\n", first, second, calc(f, s, "-"))
	fmt.Printf("%s x %s = %d\n", first, second, calc(f, s, "*"))
	fmt.Printf("%s / %s = %d\n", first, second, calc(f, s, "/"))
	
}
