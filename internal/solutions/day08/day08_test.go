package day08

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `""`, want: 2},
		{input: `"abc"`, want: 2},
		{input: `"aaa\"aaa"`, want: 3},
		{input: `"\x27"`, want: 5},
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

func TestPart2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `""`, want: 4},
		{input: `"abc"`, want: 4},
		{input: `"aaa\"aaa"`, want: 6},
		{input: `"\x27"`, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
