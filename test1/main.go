package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ForInt() int {
	OptionNumberFitst := bufio.NewReader(os.Stdin)
	OptionNumber, _ := OptionNumberFitst.ReadString('\n')
	OptionNumber = strings.TrimSpace(OptionNumber)
	Number, _ := strconv.Atoi(OptionNumber)
	return Number
}

func main() {
	for {
		fmt.Println("[0] : 종료")
		fmt.Println("[1] : 덧셈")
		fmt.Println("[2] : 뺄셈")
		fmt.Println("[3] : 곱셈")
		fmt.Println("[4] : 나눗셈")

		Number := ForInt()

		if Number == 0 {
			fmt.Println("종료 합니다")
			break
		} else if Number == 1 {
			for {
				FirstNumber := Number
				if FirstNumber == 0 {
					fmt.Println("덧셈을 종료 합니다.")
					break
				}
				SecondNUmber := Number
				Final := FirstNumber + SecondNUmber
				fmt.Printf("덧셈 결과 : %d", Final)
			}
		} else if Number == 2 {
			for {
				FirstNumber := Number
				SecondNUmber := Number
			}
		} else if Number == 3 {
			for {
				FirstNumber := Number
				SecondNUmber := Number
			}
		} else if Number == 4 {
			for {
				FirstNumber := Number
				SecondNUmber := Number
			}
		}
	}
}
