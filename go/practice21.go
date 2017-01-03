package main

import "fmt"

func main () {
	months := map[string] string {
		"1": "January",
		"2": "February",
		"3": "March",
		"4": "April",
		"5": "May",
		"6": "June",
		"7": "July",
		"8": "August",
		"9": "Semtember",
		"10": "October",
		"11": "Nobemver",
		"12": "December",
	}
	
	var month string
	fmt.Print("월을 입력 하세요~ ")
	fmt.Scanln(&month)
	
	fmt.Println(months[month])
}
