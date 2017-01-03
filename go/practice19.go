package main

/**
BMI계산기
키, 몸무게를 입력받아 체질량 지수를 계산하는 프로그램을 작성하라

bmi = (weight / (height * height))

bmi값이 18.5~25 사이로 나타나면 정상적인 몸무게라고 출력하고
그렇지 않는 경우는 과체중이나 저체중으로 나타낸다음
의사와 상의하라는 문구도 출력해보자
 */

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const lowerBound = 18.5
const higherBound = 25

func inFloat(txt string) float64 {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(txt)
		scan.Scan()
		line := scan.Text()
		f, e := strconv.ParseFloat(line, 64)
		if e != nil {
			fmt.Println(e)
			continue
		}
		return f
	}
}


func main(){
	height := inFloat("키? ")
	weight := inFloat("몸무게? ")
	bmi :=  (weight / (height*0.01 * height*0.01))
	
	fmt.Printf("%.2f bmi\n", bmi)
	
	if bmi >= lowerBound && bmi <= higherBound {
		fmt.Println("OK")
	} else if bmi < lowerBound {
		fmt.Println("저체중이야 병원가야 될것 같은데??")
	} else {
		fmt.Println("과체중이야 병원가야 될것 같은데??")
	}
}
