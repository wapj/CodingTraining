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

/**
트레이닝1 팁계산기 만들기
What is the bill? 200
What is the tip percentage? 15
The tip is $30
The total is $230
 */

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