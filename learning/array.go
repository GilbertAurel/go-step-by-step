package learning

import "fmt"

func Array() {
	var zeros [3]int
	array := []int{1, 2, 3}
	multiArray := [][]int{array, zeros[:]}

	fmt.Println(zeros)
	fmt.Println(array)

	zeros[2] = 100
	fmt.Println(zeros)
	fmt.Println(len(zeros))

	fmt.Println(multiArray)
}
