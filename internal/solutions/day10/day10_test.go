package day10

import (
	"slices"
	"testing"
)

func TestProgress(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "1", want: "11"},
		{input: "11", want: "21"},
		{input: "21", want: "1211"},
		{input: "1211", want: "111221"},
		{input: "111221", want: "312211"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := progress(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: "31133122", want: []string{"3", "1133", "122"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := split(tt.input, 1)
			if slices.Compare(got, tt.want) != 0 {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRunWithCache(t *testing.T) {
	tests := []struct {
		input      string
		iterations int
		want       int
	}{
		{input: "3113322113", iterations: 40, want: 329356},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := runWithCache(tt.input, tt.iterations)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestRunConway(t *testing.T) {
	tests := []struct {
		input      string
		iterations int
		want       int
	}{
		{input: "3113322113", iterations: 40, want: 329356},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := runConway(tt.input, tt.iterations)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
