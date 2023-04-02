package main

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	// 短い方を基準にして良い
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	res := make([][]int, 0, len(nums1)+len(nums2))
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		// 両方取り出す。先頭比較
		arr1, arr2 := nums1[i], nums2[j]
		index1, index2 := arr1[0], arr2[0]
		var index, val int
		if index1 <= index2 {
			// 左の方が小さいので、先に消費する
			index, val = index1, arr1[1]
			i++
		} else {
			// 右の方が小さいので、先に消費する
			index, val = index2, arr2[1]
			j++
		}
		// そのインデックスの要素をすでに末尾に格納済みかチェック
		resLast := len(res) - 1
		if len(res) != 0 {
			lastIndex := res[resLast][0]
			if lastIndex != index {
				// 存在しないので、追加
				res = append(res, []int{index, val})
			} else {
				// 存在するので、加算
				res[resLast][1] += val
			}
		} else {
			// そもそも一つも要素が入ってないので、追加
			res = append(res, []int{index, val})
		}

	}

	return res
}
