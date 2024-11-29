package main

import "fmt"

type student struct {
	id   int
	name string
	gpa  float64
}

func main() {
	var student student

	student.id = 202444064
	student.name = "박상민"
	student.gpa = 4.5
	fmt.Println(student.gpa)

	var student2 student
	student2.id = 202444065
	student2.name = "박상민2"
	student2.gpa = 4.49

	fmt.Println(student2.gpa)
}
