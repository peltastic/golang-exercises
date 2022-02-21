package main

import "fmt"

func main() {
	fmt.Println("Enter first no:")
	var firstno int
	fmt.Scanln(&firstno)
	fmt.Println("Enter Second no: ")
	var secondno int
	fmt.Scanln(&secondno)

	var result = firstno + secondno

	fmt.Printf("this is the ans: %d\n", result)
}
