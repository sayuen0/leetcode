package main

var m = make(map[int]int)

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	if _, ok := m[n]; !ok {
		m[n] = climbStairs(n-1) + climbStairs(n-2)
	}
	return m[n]
}
