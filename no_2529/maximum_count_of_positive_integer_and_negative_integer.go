package main

func maximumCount(nums []int) int {
	return count(nums, 0, len(nums)-1)
}

// 指定された範囲内での
func count(nums []int, left, right int) int {
	// 中央のインデックスを決定
	mid := left + (right - left) - 1

	if nums[mid] == 0 {
		// 中央値が0
		// 左右それぞれにどこまで0が続くかで分割する
	}
	panic("implement me")
}
