package main

import (
	"fmt"
)

func main() {
	//얜 배열이다
	names := [5]string{"가", "나", "다"}
	fmt.Println(names)

	//얜 slice다
	products := []string{"상자", "사과"}
	fmt.Println(products)

	// 이거 하면 오류남.. 왜냐고? 명시를 안해줬거든. slice에 추가하려면 아래 append 함수 사용함
	//products[2] = "하나추가요"
	//fmt.Println(products)

	//append 함수는 ArrayList구현체인거 같음. 왜냐면.. add하면 new Array( old array +1 뭐였지 기억이안나 쉬프트연산인데
	//이거해서 +1하고 slice를 새로 리턴해주기때문)
	products = append(products, "고릴라나스가")
	fmt.Println(products)
	fmt.Println(append(products, "고릴라나스가"))
}
