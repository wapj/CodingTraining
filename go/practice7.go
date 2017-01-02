package main

import (
	"fmt"
	"strconv"
)


/**
방의 면접을 계산하는 프로그램을 작성하라
방의 길이와 폭을 피트단위로 입력 받은 다음
제곱 피트와 제곱 미터로 면접을 나타내 보자

제곱피트에서 제곱미터롤 변환하는 식은 다음과 같다

meter = feet / 0.09290304
 */


const METER_CONST float64 = 0.09290304
func main() {
	var width, height string
	fmt.Print("방의 길이는? ")
	fmt.Scanln(&width)
	fmt.Print("방의 폭? ")
	fmt.Scanln(&height)
	
	w, _ := strconv.Atoi(width)
	h, _ := strconv.Atoi(height)
	fmt.Printf("You entered dimensions of %s feet by %s feet \n", width, height)
	fmt.Println("The area is")
	fmt.Printf("%d square feet\n", w * h)
	fmt.Printf("%f squere meters\n", float64(w * h) * METER_CONST)
}
