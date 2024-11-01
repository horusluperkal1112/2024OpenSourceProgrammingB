package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	number = strings.TrimSpace(number) // 줄바꿈, 띄엇쓰기 탭 등 제거 python strip과 유사
	n, _ := strconv.Atoi(number)       // 실수형 64bit 타입의 형변환
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	i := 2
	for i < n {
		if n%i == 0 {
			count = count + 1
		}
		i++
	}
	if count == 0 {
		fmt.Printf("%d는(은) 소수입니다.", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다.", n)
	}
}
