package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
)

/**
원금
연이율
기간
 */


func in(txt string) int {
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(txt)
		scan.Scan()
		line := scan.Text()
		v, e := strconv.Atoi(line)
		if e != nil {
			fmt.Println(e)
		} else {
			return v
		}
	}
}

func inFloat(txt string) float64{
	scan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(txt)
		scan.Scan()
		line := scan.Text()
		v, e := strconv.ParseFloat(line, 64)
		if e != nil {
			fmt.Println(e)
		} else {
			return v
		}
		
	}
}

func main() {
	principal := inFloat("원금 ?")
	interest := inFloat("연이율?")
	years := inFloat("기간?")
	
	fmt.Printf("%.f년 동안 %.2f 이자로 %.f원을 벌수 있슴다",
		years, interest, float64(principal + (float64(principal) * interest * 0.01 * float64(years))))
	
}
