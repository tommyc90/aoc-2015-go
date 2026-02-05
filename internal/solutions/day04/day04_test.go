package day04

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int64
	}{
		{input: "abcdef", want: 609043},
		{input: "pqrstuv", want: 1048970},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := part1(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
