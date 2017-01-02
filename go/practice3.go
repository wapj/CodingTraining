package main

import (
"fmt"
	"bufio"
	"os"
	"strings"
)

/**
인용구와 그 말을 한 사람을 입력 받는 프로그램을 작성하라.
인용구와 사람의 이름은 다음의 출력 예와 같이 나타내 보자

---------------
출력예
---------------
What is the quote? 인용구문
Who said it? 이름
이름 says, "인용구문"
 */

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("인용구? ")
	quote, _ := reader.ReadString('\n')
	fmt.Print("말한 사람?")
	name, _ := reader.ReadString('\n')
	fmt.Printf("%s가 말하길, \"%s\"라고 합니다", strings.TrimSpace(name), strings.TrimSpace(quote))
}
