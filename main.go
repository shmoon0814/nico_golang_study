package main

import "fmt"

func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	}
	return true
}

func canIDrinkKorea(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge == 18:
		return true
	}
	return true
}

func main() {

	fmt.Println(canIDrink(16))
	fmt.Println(canIDrinkKorea(16))
}
