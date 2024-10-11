package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("이름 입력 : ")
	r := bufio.NewReader(os.Stdin)
	name, err := r.ReadString('\n')
	fmt.Println(name)
	fmt.Println(err)
}
