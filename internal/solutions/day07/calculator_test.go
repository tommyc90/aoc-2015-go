package day07

import "testing"

func TestCalculate(t *testing.T) {
	tests := []struct {
		name         string
		instructions map[string]instruction
		targetSignal string
		want         uint16
	}{
		{
			name: "AND",
			instructions: map[string]instruction{
				"r": {operand1: "a", operand2: "b", operator: and, outputSignal: "r"},
				"a": {operand1: "2", operator: wire, outputSignal: "a"},
				"b": {operand1: "4", operator: wire, outputSignal: "b"},
			},
			targetSignal: "r",
			want:         0,
		},
		{
			name: "OR",
			instructions: map[string]instruction{
				"r": {operand1: "a", operand2: "b", operator: or, outputSignal: "r"},
				"a": {operand1: "2", operator: wire, outputSignal: "a"},
				"b": {operand1: "4", operator: wire, outputSignal: "b"},
			},
			targetSignal: "r",
			want:         6,
		},
		{
			name: "LSHIFT",
			instructions: map[string]instruction{
				"r": {operand1: "a", operand2: "2", operator: lshift, outputSignal: "r"},
				"a": {operand1: "2", operator: wire, outputSignal: "a"},
			},
			targetSignal: "r",
			want:         8,
		},
		{
			name: "RSHIFT",
			instructions: map[string]instruction{
				"r": {operand1: "a", operand2: "2", operator: rshift, outputSignal: "r"},
				"a": {operand1: "2", operator: wire, outputSignal: "a"},
			},
			targetSignal: "r",
			want:         0,
		},
		{
			name: "NOT",
			instructions: map[string]instruction{
				"r": {operand1: "a", operator: not, outputSignal: "r"},
				"a": {operand1: "0", operator: wire, outputSignal: "a"},
			},
			targetSignal: "r",
			want:         65535,
		},
		{
			name: "WIRE",
			instructions: map[string]instruction{
				"r": {operand1: "a", operator: wire, outputSignal: "r"},
				"a": {operand1: "123", operator: wire, outputSignal: "a"},
			},
			targetSignal: "r",
			want:         123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newCalculator(tt.instructions)
			got, err := g.calculate(tt.targetSignal)
			if err != nil {
				t.Fatal("unexpected error")
			}
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
