package main

func isValid(s string) bool {
	if len([]rune(s))%2 != 0 {
		return false
	}
	stack := make([]rune, 0, len([]rune(s)))
	for _, r := range s {
		switch r {
		case '(', '{', '[':
			// push
			stack = append(stack, r)
		default:
			if len(stack) == 0 {
				return false
			}
			var shouldBe rune
			switch r {
			case ')':
				shouldBe = '('
			case '}':
				shouldBe = '{'
			case ']':
				shouldBe = '['
			}

			// check if opening bracket is pair
			lastIndex := len(stack) - 1
			l := stack[lastIndex]
			if l != shouldBe {
				return false
			}

			// pop
			stack = stack[:lastIndex]
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}
