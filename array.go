package main

import "fmt"

func main() {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(num)
	var n [10]int /* n is an array of 10 integers */
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	multi_dimensional_array()
}

func multi_dimensional_array() {
	/* an array with 5 rows and 2 columns*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var i, j int
	fmt.Printf("capacity of app: %x\n", cap(a))

	/* output each array element's value */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
