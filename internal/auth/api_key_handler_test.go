// Package auth provides authentication utilities and moddleware for the server.
package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseAccountIdParam check that the accountId is correctly converted to a uint64.
func TestParseAccountIdParam(t *testing.T) {

	// Declare Tests
	tests := []struct {
		input          string
		expectedOutput uint64
		expectError    bool
	}{
		{"44433", 44433, false},
		{"54400001111", 54400001111, false},
		{"18446744073709551615", 18446744073709551615, false},
		{"18446744073709551616", 0, true}, // uint64 overflow
		{"44433a", 0, true},
		{"-44433", 0, true},
		{"thisisastring", 0, true},
		{"", 0, true},
	}

	// Run Tests
	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			output, err := parseAccountIdParam(test.input)

			assert.Equal(t, test.expectedOutput, output)

			if test.expectError {
				assert.Error(t, err)
			}
		})
	}

}
