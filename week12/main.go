package main

import (
	"fmt"
	"time"
)

func main() {
	//var dates [3]time.Time
	//dates[0] = time.Unix(0, 0)
	//dates[2] = time.Unix(1708012345, 1)

	//fmt.Println(dates[0], dates[1], dates[2])

	dates := [3]time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(1708012345, 0), // 마지막에도 ","을 작성해야 해당 방법으로 작동
	}
	//fmt.Println(dates[0], dates[1], dates[2])
	//fmt.Println(dates)
	//fmt.Printf("%#v\n", dates)
	/*
		for i := 0; i < len(dates); i++ {
			fmt.Println(1, dates[i])
		}
	*/
	for i, date := range dates { //ㅣlike python for in
		fmt.Println(i, date)
	}

}
