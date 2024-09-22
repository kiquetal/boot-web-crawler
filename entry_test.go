package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
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
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.input
			exitCode := main()
			if exitCode != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected %v, got %v", i, tc.name, tc.expected, exitColde)
			}

		})
	}

}
