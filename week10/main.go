package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isPrime(n int) bool {

	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		i := 3
		for i*i <= n {
			if n%i == 0 {
				return false
			}
			fmt.Printf("%d ", i)
			i = i + 2
		}
	}
	return true
}
func main() {
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	number = strings.TrimSpace(number) // ctrl c : 강제 종료
	n, _ := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	if isPrime(n) {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}
}
