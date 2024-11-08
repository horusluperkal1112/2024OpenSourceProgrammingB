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
	fmt.Print("첫째 정수 입력(시작값) : ")
	in := bufio.NewReader(os.Stdin)
	a, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	a = strings.TrimSpace(a) // ctrl c : 강제 종료
	n1, _ := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("둘째 정수 입력(종료값) : ")
	//in := bufio.NewReader(os.Stdin)
	b, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	b = strings.TrimSpace(b) // ctrl c : 강제 종료
	n2, _ := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
