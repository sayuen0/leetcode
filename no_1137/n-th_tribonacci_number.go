package main

func tribonacci(n int) int {
	m := make(map[int]int)
	m[0], m[1], m[2] = 0, 1, 1
	return _tribonacci(n, m)
}

func _tribonacci(n int, m map[int]int) int {
	if n <= 2 {
		return m[n]
	}
	if _, ok := m[n]; !ok {
		m[n] = _tribonacci(n-1, m) + _tribonacci(n-2, m) + _tribonacci(n-3, m)
	}

	return m[n]
}
