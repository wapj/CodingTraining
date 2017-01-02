package main

/**
연습 2 : 문자열을 입력받은 다음 입력받은 문자열의 길이와 문자열을 출력하는 프로그램을 작성하라
 */

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("What is the input string? ")
	fmt.Scanln(&str)
	fmt.Printf("%s %d ", str, len(str))
}