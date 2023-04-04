package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		t := nums[mid]
		if t == target {
			return mid
		} else if t > target {
			// 狙いの値が実際に得られた値より小さい、狙いの値はもっと左にいる
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
