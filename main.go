package main

import "fmt"

const t = "자료형이없는변수"

var aa string = "자료형이 있는 변수"

//aa = "자료형이 있는놈 변경"  안됨 func 안에서만 변경 가능 상수 const 써야 할듯 하다..
//bb := "스트링" 자료형 자동 설정도 func 밖에선 안됨
func main() {
	const name = "상혁"
	//name = "낄렵" //const는 변경 할 수 없습니다.

	var sang string = "하이"
	sang = "변경가능"

	moon := "초기화"
	moon = "go가 타입을 초기화 해줌"

	fmt.Println(name)
	fmt.Println(sang)
	fmt.Println(moon)
	fmt.Println(t)
	fmt.Println(aa)
	fmt.Println(bb)

}
