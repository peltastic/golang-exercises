package main

import "fmt"

func main() {
	/* local variable definition */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch grade {
	case "A":
		fmt.Printf("Excellent!\n")
	case "B", "C":
		fmt.Printf("Well done\n")
	case "D":
		fmt.Printf("You passed\n")
	case "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)
}
