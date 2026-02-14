package day05

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "ugknbfddgicrmopn", want: true},
		{input: "aaa", want: true},
		{input: "jchzalrnumimnmhp", want: false},
		{input: "haegwjzuvuyypxyu", want: false},
		{input: "dvszwmarrgswjxmb", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := part1(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "qjhvhtzxzqqjkmpb", want: true},
		{input: "xxyxx", want: true},
		{input: "uurcxstgmygtbstg", want: false},
		{input: "ieodomkazucvgmuy", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := part2(tt.input)
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
