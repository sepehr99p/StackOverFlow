package basics

import "fmt"

func ExpressionSwitch() {
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid day")
	}

	switch {
	case day == 1:
		fmt.Println("Monday")
	case day == 4:
		fmt.Println("Thursday")
	case day > 5:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day")
	}

}

func TypeSwitch() {
	//A Type Switch is used to branch on the type of an interface value, rather than its value.
	//This is particularly useful when dealing with variables of unknown types.
	var day interface{} = 4
	switch v := day.(type) {
	case int:
		switch v {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		default:
			fmt.Println("Invalid day")
		}
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}
