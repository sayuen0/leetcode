package main

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return binarySearch(nums, target, mid+1, right)
	}

	return binarySearch(nums, target, left, mid-1)
}
