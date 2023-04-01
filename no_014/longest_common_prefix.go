package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	cp := strs[0]
	for i, s := range strs {
		if i == 0 {
			continue
		}

		comp := commonPrefix(s, cp)
		if len(cp) > len(comp) {
			cp = comp
		}
		if cp == "" {
			break
		}
	}
	return cp
}

// 共通しているprefixを返す
func commonPrefix(a, b string) string {
	ar, br := []rune(a), []rune(b)
	// 短い方を基準にする
	if len(ar) > len(br) {
		ar, br = br, ar
	}

	cp := make([]rune, 0, len(ar))
	for i := 0; i < len(ar); i++ {
		if ar[i] != br[i] {
			break
		}
		cp = append(cp, ar[i])
	}

	return string(cp)
}
