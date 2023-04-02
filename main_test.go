package main

import (
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Reverse Linked List", "reverse_linked_list"},
		{"Maximum Count of Positive Integer and Negative Integer", "maximum_count_of_positive_integer_and_negative_integer"},
		{"Two Sum", "two_sum"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := toSnakeCase(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, actual)
			}
		})
	}
}
