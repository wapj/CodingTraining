package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

/**
세가지 물건의 가격과 수량을 각각 입력 받은 다음
소계를 구하고 소계에 대한 5.5%의 세금을 계산하자.

------------
출력예
------------
아이텤1 가격
아이템1 갯수
아이텤2 가격
아이템2 갯수
아이텤3 가격
아이템3 갯수
세금
합계금액
 */

func in(txt string) int {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(txt)
		scan.Scan()
		line := scan.Text()
		l, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		return l
	}
}

func main() {
	item1price := in("아이템 1 가격?")
	item1cnt := in("아이템 1 갯수?")
	
	item2price := in("아이템 2 가격?")
	item2cnt := in("아이템 2 갯수?")
	
	item3price := in("아이템 3 가격?")
	item3cnt := in("아이템 3 갯수?")
	
	subtotal := float64(item1cnt * item1price + item2cnt * item2price + item3cnt * item3price)
	tax := subtotal * 0.055
	total := subtotal + tax
	
	fmt.Printf("소계 : %.2f\n", subtotal)
	fmt.Printf("세금 : %.2f\n", tax)
	fmt.Printf("합계 : %.2f\n", total)
	
}
