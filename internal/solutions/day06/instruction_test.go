package day06

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  instruction
	}{
		{input: `turn on 0,0 through 999,999`, want: instruction{mode: on, fromRow: 0, toRow: 999, fromCol: 0, toCol: 999}},
		{input: `turn off 0,0 through 999,999`, want: instruction{mode: off, fromRow: 0, toRow: 999, fromCol: 0, toCol: 999}},
		{input: `toggle 0,0 through 999,999`, want: instruction{mode: toggle, fromRow: 0, toRow: 999, fromCol: 0, toCol: 999}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseInstruction(tt.input)
			if err != nil {
				t.Fatal("unexpected error")
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse_InvalidInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		{input: `turn onn 0,0 through 999,999`},
		{input: `turn on 0,abc through 999,999`},
		{input: `turn on 0,0 through -999,999`},
		{input: `turn on 999,0 through 0,999`},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := parseInstruction(tt.input)
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
