package day03

import "testing"

func TestPart1(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `>`, want: 2},
		{input: `^>v<`, want: 4},
		{input: `^v^v^v^v^v`, want: 2},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPart1_InvalidInput(t *testing.T) {
	_, err := part1(">_")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{input: `^v`, want: 3},
		{input: `^>v<`, want: 3},
		{input: `^v^v^v^v^v`, want: 11},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestPart2_InvalidInput(t *testing.T) {
	_, err := part2(">_")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
