package main

import (
	"fmt"
	"reflect"
)

func main() {

	i := 13
	var f float64 = 12.9
	c1 := 'Z'
	c2 := '김'
	fmt.Printf("value i : %d, value f : %f\n", i, f)
	//fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f) // 변환

	fmt.Printf("%d * %f = %d\n", i, f, i*int(f))
	fmt.Println(c1, c2)    // 10진수 출력
	fmt.Printf("%x\n", c2) // 유니코드 '김'에 대한 16진수 출력, home.unicode.org(유니코드 사이트)
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(c1), reflect.TypeOf(c2))
}
