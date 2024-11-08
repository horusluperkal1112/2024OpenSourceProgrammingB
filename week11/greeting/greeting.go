package greeting

import "fmt"

// 소문자로 시작하는 함수는 외부 접근이 불가하나 내부에서는 접근이 가능하다
// 대문자 함수로 우회할 경우 외부에서도 접근할 수 있다.
func hi(name string) {
	fmt.Printf("Hi %s!\n", name)
}
func hello(name string) {
	fmt.Printf("Hellow %s~\n", name)
}

func EnglishGreetings(name string, name2 string) {
	hello(name)
	hi(name2)
}
