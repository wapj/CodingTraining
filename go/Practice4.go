package main

/**
Mad Libs
 
 명사 동사 형용사 부사를 받아서 문장을 만들어 보자~
 */


import "fmt"

func main(){
	
	var noun, verb, adj, adv string
	fmt.Print("명사 : ")
	fmt.Scanln(&noun)
	fmt.Print("동사 : ")
	fmt.Scanln(&verb)
	fmt.Print("형용사 : ")
	fmt.Scanln(&adj)
	fmt.Print("부사 : ")
	fmt.Scanln(&adv)
	
	fmt.Printf("%s %s %s %s", adj, noun, adv, verb)
	
}
