package main

import (
	"fmt"
	"time"
	"strconv"
)

/**
퇴직 계산기
나이, 정년을 입력 받은후 몇년도에 퇴직하게 되는지 알려줌
 */


func getYear() int{
	year, _, _ := time.Now().Date()
	return year
}

func main() {
	var age, retire string
	fmt.Print("나이? ")
	fmt.Scanln(&age)
	fmt.Print("정년? ")
	fmt.Scanln(&retire)
	a, _ := strconv.Atoi(age)
	r, _ := strconv.Atoi(retire)
	
	fmt.Printf("%dyear, so you can retire in %d", getYear(), getYear() + r - a)
}
