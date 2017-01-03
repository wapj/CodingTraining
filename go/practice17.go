package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

/**
혈중 알코올 농도 게산하기
Blood Alcohol Content(BAC)
BAC = (A x 5.14 / W x r ) - 0.15 x H

A : 총 알콜 소비량 (온스 단위)
W : 몸무게 (파운드)
r : 성별에 따른 알코올 흡수비 갯수
  - 0.73 : 남자
  - 0.6 : 여자
H : 술을 마신후 경과 시간

법적으로 허용되는 수치는 0.08 미만

---------------------
출력예시
---------------------
알코올 소비량? 20
몸무게? 30
술마시고 경과시간? 4
남자 or 여자??M
혈중 알코올 농도 1.901
---------------------
 */

func inFloat(txt string) float64 {
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

func calcBAC(a, w, h float64, g string) float64 {
	var r float64
	if g == "M" {
		r = 0.73
	} else {
		r = 0.6
	}
	return (a * 5.14 / w * r) - 0.15 * h
}

func main(){
	a := inFloat("알코올 소비량? ")
	w := inFloat("몸무게? ")
	h := inFloat("술마시고 경과시간? ")
	
	var g string
	fmt.Print("남자 or 여자??")
	fmt.Scanln(&g)
	
	fmt.Printf("혈중 알코올 농도 %.3f", calcBAC(a, w, h ,g))
	
}
