package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

/**
피자를 정확하게 나누는 프로그램을 작성하라
사람 수
피자 갯수
조각 갯수를 입력 받는데 이때 조각 갯수는 짝수여야 한다.
 */

func in(enter string) int {
	var line string
	var l int = 0
	scan := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print(enter)
		scan.Scan()
		line = scan.Text()
		l, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
		} else {
			return l
			break
		}
	}
	return l
}


func main() {
	people := in("몇명이요? ")
	pizza := in("피자 몇판? ")
	peices := in("한판에 몇 조각? ")
	fmt.Printf("한사람당 %d 조각 먹을 수 있다\n", pizza * peices / people)
	fmt.Printf("%d 조각 남았다\n", pizza * peices % people)
}
