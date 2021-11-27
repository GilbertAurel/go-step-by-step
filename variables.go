package main

import (
	"fmt"
	"math"
)

func main() {
	var name string = "golang"

	var number1, number2 int8 = 7, 2
	var result float32 = float32(number1) / float32(number2)

	var fruit1 string = "apple"
	fruit2 := "orange"

	const pi = 22.0 / 7.0
	const hundred = 3e2
	const alotOfZero = 3e10

	fmt.Println(name)

	fmt.Println("1+1 =", 1+1)
	fmt.Printf("1+1 = %v \n", 1+1)

	fmt.Println(result)

	fmt.Println(true)
	fmt.Println(number1 < number2)
	fmt.Println(!(false || true))

	fmt.Println(fruit1, fruit2)

	fmt.Println(pi)
	fmt.Println(math.Sin(hundred))
	fmt.Println(int64(alotOfZero))
}
