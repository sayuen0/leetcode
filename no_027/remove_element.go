package main

func removeElement(nums []int, val int) int {
	var l, m int
	for i := 0; i < len(nums); i++ {
		if val != nums[i] {
			nums[l] = nums[i]
			l++
		} else {
			m++
		}
	}

	return l
}
