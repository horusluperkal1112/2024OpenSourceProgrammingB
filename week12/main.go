package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a slice by slicing make functuon

	var gpaSlice []float64
	gpaSlice = make([]float64, 3)
	gpaSlice[0] = 4.1
	gpaSlice[1] = 4.5
	gpaSlice[2] = 3.5
	fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	// Create a slice by slicing literal
	//gpas := [5]float32{3.5, 4.1, 4.5, 3.9, 4.23}
	//fmt.Println(gpas, reflect.TypeOf(gpas))

	//gpaSlice := gpas[1:4]
	//fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	//gpaSlice := gpas[1:4]
	//fmt.Println(gpaSlice, reflect.TypeOf(gpaSlice))

	//gpas := [5]float32{3.5, 4.1, 4.5, 3.9, 4.23}
	//fmt.Println(gpas, reflect.TypeOf(gpas))

}
