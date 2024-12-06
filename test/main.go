package main

import "fmt"

func main() {
	// var array [5]int  배열 기본 (int는 자료형)
	fmt.Println("최고점, 평균점, 총 점수 개수 구하기")
	scores := [5]int{87, 92, 78, 90, 88}

	// 최고점 구하기
	high := scores[0]
	for _, score := range scores {
		if score > high {
			high = score
		}
	}
	// 평균 구하기
	sum := 0
	for _, score := range scores {
		sum += score
	}
	average := float64(sum) / float64(len(scores))
	// 총점 구하기

	fmt.Printf("최고점은 %d점 입니다\n", high)
	fmt.Printf("총점은 %d점 입니다\n", sum)
	fmt.Printf("평균점은 %.2f점 입니다", average)
}
