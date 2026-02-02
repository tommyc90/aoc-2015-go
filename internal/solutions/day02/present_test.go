package day02

import "testing"

func TestPresent_InvalidInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		{input: `1x1`},
		{input: `1x1x1x1`},
		{input: `abc`},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := parsePresent(tt.input)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
