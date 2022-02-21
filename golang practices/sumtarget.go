package main

import "fmt"

func main() {
	result := twoSum([]int{4, 2, 6}, 8)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var index_one, index_two int
LOOP:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			println(i)
			fmt.Println(i + j)
			sum := nums[i] + nums[j]
			fmt.Printf("loop %v result %v \n", i, sum)
			if sum == target {
				index_one = i
				index_two = j
				break LOOP
			}

		}
	}

	return []int{index_one, index_two}

}
