package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

/**
천장을 칠하는데 필요한 페인트 량을 구하는 프로그램을 작성하라
길이와 폭을 입력 받은 다음 1리터에 9m^2를 칠한다고 가정하여 계산하자.
그리고 천장을 칠하는데 필요한 페인트의 양을 정수로 표현해 보자
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
	width := in("가로? ")
	height := in("세료? ")
	
	area := width * height
	cnt := math.Ceil(float64(area) / 9)
	
	fmt.Printf("면적? %d 제곱미터 \n", area)
	fmt.Printf("페인트 %d개 사면됩니다\n", int(cnt))
	
	
}
