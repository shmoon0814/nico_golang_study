package main

import (
	"fmt"
	"strings"
)

//naked return 음 글쌔 얘는... 잘... 흠... 리턴에 변수명, 자료형 선언해주면 내부에서 로직 돌린 후 리턴만하면 자동으로 넘긴다.. 뭐 이런
func checkLengthAndWord(words string) (length int, uppercase string) {
	length = 4
	uppercase = strings.ToUpper(words)
	return
}

//defer 시벌 이거 대박 promise...
func deferCheck() {
	defer fmt.Println("이게 defer 야.. 특정 동작행위를 끝내고 마지막에 동작한다")
	fmt.Println("얘가 먼저 실행되고")
}

func main() {
	fmt.Println(checkLengthAndWord("sanghyuk"))
	deferCheck()
}
