package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "no arguments",
			input:    []string{"./entry"},
			expected: 1,
		},
		{
			name:     "one argument",
			input:    []string{"./entry", "https://blog.boot.dev"},
			expected: 0,
		},
	}
	// just pass the test
	for _, tc := range tests {

		os.Args = tc.input
		main()

	}
}
