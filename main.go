package main

import "fmt"

func main() {
	a := 2
	b := a
	fmt.Println(&a, &b)
	//주소 확인

	c := 4
	d := &c //주소값을 저장하는 법
	fmt.Println(&c, d)
	fmt.Println(*d) //여기서 d가 왜 4가 나오냐면
	//자동타입 추론으로 인하여 &가 달렸으면 주소를 나타내는 자료형이라서
	//d는 메모리 주소를 나타내는 자료형이고 *는 메모리 자료형의 주소의 값을 나타내는 것

	e := 3
	f := &e
	*f = 20
	fmt.Println(e, &e)
	//주소값을 가지고 해당 값을 변경할 수도 있다.

}
