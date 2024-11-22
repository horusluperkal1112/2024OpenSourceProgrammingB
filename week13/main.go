package main

import (
	"fmt"
)

func main() {

	gpas := [5]float32{3.5, 4.1, 4.5, 3.9, 4.23}
	gpas_slice := gpas[1:4]
	//gpas_slice[1] = 2.71
	gpas[1] = 2.71
	//gpas_slice = append(gpas_slice, 4.3)
	gpas_slice = append(gpas_slice, 4.3, 5.55)
	fmt.Println(len(gpas_slice), gpas_slice, gpas)

}
