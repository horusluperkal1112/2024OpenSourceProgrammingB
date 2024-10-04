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
	fmt.Println(c1, c2)
	fmt.Println(reflect.TypeOf(f), reflect.Type(i))
}
