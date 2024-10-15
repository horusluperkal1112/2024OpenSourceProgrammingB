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
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}
	score = strings.TrimSpace(score)              // 줄바꿈, 띄엇쓰기 탭 등 제거 python strip과 유사
	realScore, _ := strconv.ParseFloat(score, 64) // 실수형 64bit 타입의 형변환

	if realScore >= 60 {
		fmt.Println("합격")
	} else {
		fmt.Println("탈락")
	}
}
