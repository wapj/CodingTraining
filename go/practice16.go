package main

import (
	"fmt"
	"strconv"
	"os"
)

func main(){
	fmt.Print("나이가 우째 되노? ")
	var a string
	fmt.Scanln(&a)
	age, err := strconv.Atoi(a)
	
	if err != nil || age < 1 {
		fmt.Println("0보다 큰 숫자를 적으세여~")
		os.Exit(1)
	}
	var str string
	
	fmt.Println(age)
	
	if age > 17 {
		str = "운전가능"
	} else {
		str = "운전안대"
	}
	fmt.Println(str)
}
