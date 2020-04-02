package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico")
	go sexyCount("sanghyuk")
	go sexyCount("gogo")
	// 이게 둘다 go 를 달면 안되는 이유가 있다.
	// go의 메인함수가 진행되는 동안에만 goroutine 이 걸림
	//즉 둘다 비동기로 실행되버리면 메인함수가?? 끝나버리지... ㅇㅋ??

	time.Sleep(time.Second * 5)
	//자 main 함수는 함수의 결과가 저장되는 곳이다..
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
