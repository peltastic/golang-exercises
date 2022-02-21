package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// reverseArr([]int{1, 2, 3})
	result := addTwoNumbers([]int{2, 4, 3}, []int{5, 6, 4})

	fmt.Println(result)
}

func addTwoNumbers(l1, l2 []int) []int {
	reverselist_1 := reverseArr(l1)
	reverselist_2 := reverseArr(l2)
	reverselist_string1 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reverselist_1)), ""), "[]")
	reverselist_string2 := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reverselist_2)), ""), "[]")
	reverselist_1_int, err1 := strconv.Atoi(reverselist_string1)
	reverselist_2_int, err2 := strconv.Atoi(reverselist_string2)
	if err2 != nil || err1 != nil {
		fmt.Println(err1, err2)
	}

	result_int := reverselist_1_int + reverselist_2_int
	result_string := strconv.Itoa(result_int)
	result_reverseStr := strings.Split(result_string, "")
	var result_reverseInt = []int{}

	for _, i := range result_reverseStr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		result_reverseInt = append(result_reverseInt, j)
	}
	result := reverseArr(result_reverseInt)

	return result

}

func reverseArr(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
