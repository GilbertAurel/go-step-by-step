package main

import (
	"fmt"
)

func main() {
	var name string = "golang"

	var number1, number2 int8 = 7, 2
	var result float32 = float32(number1) / float32(number2)

	var fruit1 string = "apple"
	fruit2 := "orange"

	fmt.Println(name)

	fmt.Println("1+1 =", 1+1)
	fmt.Printf("1+1 = %v \n", 1+1)

	fmt.Println(result)

	fmt.Println(true)
	fmt.Println(number1 < number2)
	fmt.Println(!(false || true))

	fmt.Println(fruit1, fruit2)
}
