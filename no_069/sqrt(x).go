package main

func mySqrt(x int) int {
	for i := 1; i <= x; i++ {
		ii := i * i
		if ii == x {
			return i
		} else if ii > x {
			return i - 1
		}
	}

	return 0
}
