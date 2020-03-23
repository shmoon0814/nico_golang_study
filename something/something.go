package something

import "fmt"

func helloWolrd() {
	fmt.Println("Hello world")
}

func SayHello() {
	fmt.Println("Say Hello")
	helloWolrd()
}
