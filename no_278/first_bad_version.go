package main

func firstBadVersion(n int) int {
	l, r := 0, n
	for mid := (l + r) / 2; l < r; mid = (l + r) / 2 {
		if isBadVersion(mid) {
			l = mid
		} else {
			r = mid + 1
		}
	}
	return l
}
