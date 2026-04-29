package day07

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  instruction
	}{
		{input: `abc AND def -> xyz`, want: instruction{operator: and, operand1: "abc", operand2: "def", outputSignal: "xyz"}},
		{input: `abc OR def -> xyz`, want: instruction{operator: or, operand1: "abc", operand2: "def", outputSignal: "xyz"}},
		{input: `abc LSHIFT 123 -> xyz`, want: instruction{operator: lshift, operand1: "abc", operand2: "123", outputSignal: "xyz"}},
		{input: `abc RSHIFT 123 -> xyz`, want: instruction{operator: rshift, operand1: "abc", operand2: "123", outputSignal: "xyz"}},
		{input: `NOT abc -> xyz`, want: instruction{operator: not, operand1: "abc", outputSignal: "xyz"}},
		{input: `123 -> xyz`, want: instruction{operator: wire, operand1: "123", outputSignal: "xyz"}},
		{input: `abc -> xyz`, want: instruction{operator: wire, operand1: "abc", outputSignal: "xyz"}},
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
		{input: `abc XOR def -> xyz`},
		{input: `abc LSHIFT def -> xyz`},
		{input: `abc NOT def -> xyz`},
		{input: `123 -> 456`},
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
