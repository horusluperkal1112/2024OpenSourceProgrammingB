package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //유닉스 시간 : 1970년 언젠가부터 지금의 초
	answer := rand.Intn(6) + 1   // 주사위 1 ~ 6 사이의 수
	// fmt.Println(answer)

	win := false

	for guesses := 0; guesses < 3; guesses++ {
		fmt.Printf("%d번의 기회가 남았습니다. 숫자 입력 : ", 3-guesses)
		in := bufio.NewReader(os.Stdin)
		input, err := in.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input) // 줄바꿈, 띄어쓰기 탭 등 제거 python strip과 유사
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(guess)

		if answer == guess {
			win = true
			fmt.Println("정답입니다.")
			break
		} else if answer > guess {
			fmt.Println("입력하신 값은 정답보다 작습니다.")
		} else {
			fmt.Println("입력하신 값은 정답보다 큽니다.")
		}

	}

	if win {
		fmt.Println("당신이 이겼습니다.")
	} else {
		fmt.Printf("당신이 졌습니다. 정답은 %d입니다\n", answer)
	}
}
