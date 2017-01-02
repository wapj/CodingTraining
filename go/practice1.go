package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
연습 1. 이름을 받아서 인사말을 출력하는 프로그램을 작성해보자
 */

func  main(){
	fmt.Print("이름이 뭐예요? ")
	var name string
	rand.Seed(time.Now().UTC().UnixNano())
	answers := []string {
		"%s 반가워요~",
		"%s 오신걸 축하해요",
		"%s 커피한잔 하면서 이야기 할래요?",
	}
	fmt.Scanln(&name)
	fmt.Printf(answers[rand.Intn(len(answers))], name)
}
