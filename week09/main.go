package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //유닉스 시간 : 1970년 언젠가부터 지금의 초
	answer := rand.Intn(6) + 1   // 주사위 1 ~ 6 사이의 수
	fmt.Println(answer)
}
