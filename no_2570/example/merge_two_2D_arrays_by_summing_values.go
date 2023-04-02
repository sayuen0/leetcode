package main

func mergeArrays(ns1 [][]int, ns2 [][]int) [][]int {
	n1, n2 := len(ns1), len(ns2)
	idx1, idx2 := 0, 0

	rst := make([][]int, 0, (n1+n2)<<1)
	for idx1 < n1 && idx2 < n2 {
		if ns1[idx1][0] == ns2[idx2][0] {
			ns1[idx1][1] += ns2[idx2][1]
			rst = append(rst, ns1[idx1])
			idx1++
			idx2++
		} else if ns1[idx1][0] < ns2[idx2][0] {
			rst = append(rst, ns1[idx1])
			idx1++
		} else {
			rst = append(rst, ns2[idx2])
			idx2++
		}
	}

	if idx2 != n2 {
		rst = append(rst, ns2[idx2:]...)
	}

	if idx1 != n1 {
		rst = append(rst, ns1[idx1:]...)
	}

	return rst
}
