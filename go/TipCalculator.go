package main

import (
	"fmt"
	"strconv"
)

func toInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		fmt.Println(e)
	}
	return i
}

func main() {
	fmt.Print("What is the bill? ")
	var bill, tipPercent string
	fmt.Scanln(&bill)
	fmt.Print("What is the tip percentage? ")
	fmt.Scanln(&tipPercent)
	tip := toInt(bill) * toInt(tipPercent) / 100
	total := toInt(bill) + tip
	fmt.Printf("The tip is $%d\n", tip)
	fmt.Printf("The total is $%d\n", total)
	
}