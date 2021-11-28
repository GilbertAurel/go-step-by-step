package learning

import "fmt"

func ClassicSwitch(menu int) {

	switch menu {
	case 1:
		fmt.Println("User pick menu 1")
	case 2, 3:
		fmt.Println("User pick menu 2 or 3")
	default:
		fmt.Println("User pick other menu")
	}
}

func SwitchType(input interface{}) {
	switch input.(type) {
	case bool:
		fmt.Println("Type is boolean")
	case int:
		fmt.Println("Type is integer")
	default:
		fmt.Println("Type is not boolean or integer")
	}
}
