package main

import (
	"fmt"
	"github.com/shmoon0814/learngo/something"
)

func main() {
	fmt.Println("Hello World!")
	something.SayHello()

	// Println 이 대문자로 실행되는 이유는
	// go 에선 Uppercase method 는 public
	// lowercase method 는 private
	// main에서는 helloWorld를 찾을 수 없지만 something 패키지 내부에서는 호출 할 수 있다.
	// 대부분의 호출 되는 메소드 들은 대문자로 시작 할 것
}
