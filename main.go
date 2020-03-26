package main

import "fmt"

func canIDrink(age int) bool {

	//일반적으로 위에 플래그 선언하고 쓰는거랑 if문 안에 넣을 수 있는거랑 차이가 없다.
	//하지만 이렇게 쓰는 이유는 해당 변수는 해당 if 문에서만 쓰기 위해서 선언한 변수다.. 라는 약속같은거..
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func main() {

	flag := canIDrink(16)
	fmt.Println(flag)

}
