package main

func maximumCount(nums []int) int {
	//find last negative number
	l, r := 0, len(nums)-1
	neg := 0
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] < 0 && (mid+1 == len(nums) || nums[mid+1] > -1) {
			neg = mid + 1 //add 1 because arrays are 0 indexed
			break
		} else {
			if nums[mid] > -1 {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	//find first positive number
	l, r = 0, len(nums)-1
	pos := 0
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > 0 && (mid-1 == -1 || nums[mid-1] < 1) {
			pos = len(nums) - mid
			break
		} else {
			if nums[mid] < 1 {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	if neg > pos {
		return neg
	}
	return pos
}
