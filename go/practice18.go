package main

/**
섭씨와 화씨를 서로 변환하는 프로그램을 작성하라
 */

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func in(txt string) float64 {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(txt)
		scan.Scan()
		line := scan.Text()
		l, e := strconv.ParseFloat(line, 64)
		
		if e != nil {
			fmt.Println(e)
			continue
		}
		return l
	}
}

func transTemperature(fc string, temp float64) float64{
	
	if fc == "C" {
		temp =(temp - 32) * 5/9
	} else {
		temp = (temp * 9*5) + 32
	}
	return temp
}

func main() {
	fmt.Print("화씨인지 섭씨인지 고르세여~ ( F or C ) ")
	var fc string
	fmt.Scanln(&fc)
	temp := in("온도 ? ")
	
	fmt.Println(transTemperature(fc, temp))
}
