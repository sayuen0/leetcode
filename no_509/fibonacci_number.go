package main

func fib(n int) int {
	m := make(map[int]int)
	m[0], m[1], m[2] = 0, 1, 1
	return _fib(n, m)
}

func _fib(n int, m map[int]int) int {
	if n <= 1 {
		return n
	}
	if _, ok := m[n]; !ok {
		m[n] = _fib(n-1, m) + _fib(n-2, m)
	}
	return m[n]
}
