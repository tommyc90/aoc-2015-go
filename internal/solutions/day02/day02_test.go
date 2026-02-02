package day02

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `2x3x4`, want: 58},
		{input: `1x1x10`, want: 43},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p, err := parsePresent(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := part1([]present{p})
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
		{input: `2x3x4`, want: 34},
		{input: `1x1x10`, want: 14},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p, err := parsePresent(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := part2([]present{p})
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
