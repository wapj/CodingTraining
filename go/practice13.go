package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

/**
복리를 통해 투자 수익을 계산하는 프로그램을 작성하라.
프로그램은 원금, 투자기간, 연이율, 연간 수익이 지급되는 횟수를 입력 받는다

공식
A = P(1 + r/n)^nt

P : 원금
r : 연이율
t : 투자기간
n : 연간 이자 지급 횟수
A : 원리금

------------
출력예시
------------
원금 얼마? 1500
연이율? 4.3
투자기간? 6
일년에 이자가 몇번 지급 되나? 4
원금 1500 연이율 4.30 6년 투자 하고 이자가 1년에 4번 나오는데 만기되면 원리금은 1938.84원이 된다

 */

func in(txt string) int {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(txt)
		scan.Scan()
		line := scan.Text()
		l, e := strconv.Atoi(line)
		if e != nil {
			fmt.Println(e)
			continue
		}
		return l
	}
}

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

func compoundInterest(p int, r float64, t int, n int) float64 {
	_p := float64(p)
	_t := float64(t)
	_n := float64(n)
	nt := _n * _t
	return _p * math.Pow(1 + r/_n,nt)
}

func main() {
	P := in("원금 얼마? ")
	r := inFloat("연이율? ")
	t := in("투자기간? ")
	n := in("일년에 이자가 몇번 지급 되나? ")
	fmt.Printf("원금 %d 연이율 %.2f %d년 투자 하고 이자가 1년에 %d번 나오는데 만기되면 원리금은 %.2f원이 된다\n",
		P, r, t, n, compoundInterest(P, r * 0.01, t, n))
}
