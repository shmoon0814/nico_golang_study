package main

import (
	"fmt"
	"strings"
)

func multiply_1(a, b int) int { //받는 상수가 같은자료형이면 이렇게 묶어라.
	return a * b
}

func multiply_2(a int, b int) int { //받는 변수가 다르면 따로선언하시던가요
	return a * b
}

//리턴을 여러개 할 수도 있다. 이거 좀 좋냉...
func checkLengthAndWord(words string) (int, string) {
	return len(words), strings.ToUpper(words)
}

func sayWord(words ...string) {
	fmt.Println(words)
}

//이건 설명이 안나와있는디.. 내가 배열 받으니 배열 넘기는거겠
func returnWord(words ...string) int {
	return len(words)
}

func main() {

	fmt.Println(multiply_1(4, 2))
	fmt.Println(multiply_2(13, 4))
	fmt.Println(checkLengthAndWord("상혁아..."))

	//리턴 두개 받는법
	leng, word := checkLengthAndWord("상혁아")
	fmt.Println(leng)
	fmt.Println(word)

	//두개중에 하나만 받고싶을때. _ 로 하면 변수 블랭크 컴파일러가 무시한다
	leng_2, _ := checkLengthAndWord("상혁아")
	fmt.Println(leng_2)

	_, word_2 := checkLengthAndWord("가나다라")
	fmt.Println(word_2)

	//for문 돌릴 꺼 이렇게 변수 ...로 선언해 놓으면 한기능 반복처리가 가능하다.
	//아 이거 혁명인디?? 나중에 포문을 좀 봐야겠는데.. 지
	//뭐 결국 파라미터 인자값을 넘기는걸 좀더 편히 할 수 있다 정도다.
	//특정 연산 처리 후 넘기는 자료형 정리가 된다면 좋긴 하겠다.
	sayWord("하나")
	sayWord("하나", "둘", "셋", "넷")

	arrayLeng := returnWord("하나", "둘", "셋", "넷", "다섯")
	fmt.Println(arrayLeng)
}
