package main

import (
	"fmt"
)

func main() {
	var name string = "golang"
	var number1 int8 = 7
	var number2 int8 = 3
	var result float32 = float32(number1) / float32(number2)

	fmt.Println(name)

	fmt.Println("1+1 =", 1+1)
	fmt.Printf("1+1 = %v \n", 1+1)

	fmt.Println(result)

	fmt.Println(true)
	fmt.Println(number1 < number2)
	fmt.Println(!(false || true))
}
