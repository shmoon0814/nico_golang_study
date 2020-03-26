package main

import "fmt"

//range 를 이용한 루르
func addSuperIndexValue(numbers ...int) (returnvalue int) {
	//인덱스를 안쓸 경우는 _ 로 무시해버리기
	for _, value := range numbers {
		returnvalue += value
	}
	return //naked return
}

func printIndexValue(numbers ...int) {
	for i, v := range numbers {
		fmt.Println(i, v)
	}
}

func oldFor(numbers ...int) {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}

func main() {
	printIndexValue(6, 5, 3, 2, 1, 0)

	fmt.Println("줄바꿈")
	value := addSuperIndexValue(1, 4, 4, 3, 2, 5, 43)
	fmt.Println(value)
	fmt.Println("줄바꿈")
	oldFor(4, 3, 2, 5, 3, 2, 1)
}
