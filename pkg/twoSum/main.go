package twoSum

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
	}
	return result
}

func Ans() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
