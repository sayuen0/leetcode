package main

func twoSum(nums []int, target int) []int {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
