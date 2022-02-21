package main

import "fmt"

func main() {

	fmt.Printf("Which of the following operations would you like to pereform: \n")
	fmt.Printf("Multiplication: \n")
	fmt.Printf("Division: \n")
	fmt.Printf("Substraction: \n")
	fmt.Printf("Addition: \n")
	fmt.Println("Enter your choice: ")

	var user_choice string
	fmt.Scanln(&user_choice)
	var answer float32
	switch user_choice {
	case "multiplication":
		answer = multiplication()
		fmt.Printf("Result: %f\n", answer)
		break
	case "division":
		answer = division()
		fmt.Printf("Result: %f\n", answer)

	}
}

func multiplication() float32 {
	fmt.Println("Enter First No: ")
	var no_1 float32
	fmt.Scanln(no_1)
	fmt.Println("Enter Second No: ")
	var no_2 float32
	fmt.Scanln(no_2)

	var result float32 = no_1 * no_2

	return result
}
func addition() float32 {
	fmt.Println("Enter First No: ")
	var no_1 float32
	fmt.Scanln(no_1)
	fmt.Println("Enter Second No: ")
	var no_2 float32
	fmt.Scanln(no_2)

	var result float32 = no_1 + no_2

	return result
}
func substraction() float32 {
	fmt.Println("Enter First No: ")
	var no_1 float32
	fmt.Scanln(no_1)
	fmt.Println("Enter Second No: ")
	var no_2 float32
	fmt.Scanln(no_2)

	var result float32 = no_1 - no_2

	return result
}
func division() float32 {
	fmt.Println("Enter First No: ")
	var no_1 float32
	fmt.Scanln(no_1)
	fmt.Println("Enter Second No: ")
	var no_2 float32
	fmt.Scanln(no_2)

	var result float32 = no_1 / no_2

	return result
}
