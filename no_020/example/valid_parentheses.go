package main

import "container/list"

func isValid(s string) bool {
	l := list.New()
	for _, v := range s {
		if v == '{' || v == '(' || v == '[' {
			l.PushFront(v)
		} else if l.Len() > 0 {
			e := l.Front()
			if v == '}' && e.Value != '{' {
				return false
			}
			if v == ')' && e.Value != '(' {
				return false
			}
			if v == ']' && e.Value != '[' {
				return false
			}
			l.Remove(e)
		} else {
			return false
		}
	}
	if l.Len() > 0 {
		return false
	}
	return true
}
