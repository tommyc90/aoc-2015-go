package day06

import "testing"

func TestApply(t *testing.T) {
	tests := []struct {
		name  string
		input instruction
		want  int
	}{
		{name: "All on", input: instruction{mode: on, fromRow: 0, toRow: 999, fromCol: 0, toCol: 999}, want: 1000000},
		{name: "One row on", input: instruction{mode: on, fromRow: 0, toRow: 999, fromCol: 0, toCol: 0}, want: 1000},
		{name: "All on by toggle", input: instruction{mode: toggle, fromRow: 0, toRow: 999, fromCol: 0, toCol: 999}, want: 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGrid(1000, 1000)
			g.applyInstruction(tt.input)
			got := g.countLights()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApply_multi(t *testing.T) {
	g := newGrid(10, 10)
	g.applyInstruction(instruction{mode: on, fromRow: 0, toRow: 9, fromCol: 0, toCol: 0})
	g.applyInstruction(instruction{mode: on, fromRow: 0, toRow: 9, fromCol: 0, toCol: 0})
	g.applyInstruction(instruction{mode: off, fromRow: 0, toRow: 9, fromCol: 0, toCol: 0})
	g.applyInstruction(instruction{mode: toggle, fromRow: 0, toRow: 9, fromCol: 0, toCol: 0})
	want := 10
	got := g.countLights()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
