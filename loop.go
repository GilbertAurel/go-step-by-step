package main

import "fmt"

func main() {

	// similar to while
	i := 1
	for i <= 3 {
		fmt.Printf("%v ", i)
		i += 1
	}

	// classic
	for j := 4; j <= 10; j++ {

		if j%2 == 0 {
			continue
		}

		fmt.Printf("%v ", j)
	}
	fmt.Print("\n")

	// without condition, until
	k := 1
	for {
		if k <= 10 {
			fmt.Printf("%v ", k)
			k++
			continue
		}

		break
	}
}
