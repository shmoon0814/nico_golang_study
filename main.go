package main

import (
	"fmt"
)

func main() {
	//선언이 신기하내.. [키 자료형] 베률
	sanghyuk := map[string]string{"name": "상혁", "age": "12"}

	//맵 추가는 이런식으로 해주면 됨
	sanghyuk["op.gg"] = "http://localhost:8080"
	//문법이 죄다 통일되고있냉.. 좋다.

	for key, value := range sanghyuk {
		fmt.Println(key, value)
	}

	for _, value := range sanghyuk {
		fmt.Println(value)
	}

	//이런 식으로 하는건.. 별로 안좋은데 이게 다 보여주냉??
	favFood := []string{"kimchi", "sushi"}

	//IDE 상에선 잘 보이는디
	//소스코드상에선 아래걸로 명시해주는게 훨씬 보기 좋다.
	nico := person{"sanghyuk", 18, favFood}

	nico2 := person{name: "ninicoco", age: 41, favFood: favFood}
	fmt.Println(nico)
	fmt.Println(nico2)
}

type person struct {
	name    string
	age     int
	favFood []string
}
