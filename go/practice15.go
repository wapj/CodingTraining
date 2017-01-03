package main

import "fmt"

const password = "abc$123"

func main() {
	fmt.Print("패스워드를 입력해주소 :")
	var pass string
	fmt.Scanln(&pass)
	if pass == password {
		fmt.Println("통과~")
	} else {
		fmt.Println("다시 시도하세요")
	}
}
