package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int
	// i = 55
	var f float64 = 1.92
	//var i int = 55
	//f := 1.92
	i := "55"
	//C언어 double : float64
	//var f float32 = 1.92
	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("%f %f\n", f, math.Ceil(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("i is", f)
	fmt.Println("i is", i)
	fmt.Print("i is", i, "\n")
	fmt.Println("i is", i)
	//fmt.Printf("i i %d\n", i)
	fmt.Printf("i i %s\n", i)

}
