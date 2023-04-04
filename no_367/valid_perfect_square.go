package main

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l, r := 0, num
	for l < r {
		mid := l + (r-l)/2
		midmid := mid * mid
		if midmid == num {
			return true
		} else if midmid > num {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return false
}
