package main

import (
	"fmt"
	"github.com/shmoon0814/learngo/mydict"
)

func main() {
	dictionay := mydict.Dictionary{}
	dictionay["상혁"] = "상혁"
	value, error := dictionay.Search("정혁")

	if error == nil {
		fmt.Println(value)
	} else {
		fmt.Println(error)
	}

	word := "hello"
	def := "Greeting"
	err := dictionay.Add(word, def)

	if err != nil {
		fmt.Println(err)
	}
	dict, errr := dictionay.Search("hello")
	fmt.Println(dict, errr)
	//=======

}
