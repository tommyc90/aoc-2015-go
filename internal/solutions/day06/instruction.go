package day06

import (
	"fmt"
	"regexp"
	"strconv"
)

type modeType int

const (
	on modeType = iota
	off
	toggle
)

type instruction struct {
	mode                           modeType
	fromRow, toRow, fromCol, toCol int
}

var instructionRe = regexp.MustCompile(`^(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)$`)

func parseInstruction(value string) (instruction, error) {
	matches := instructionRe.FindStringSubmatch(value)
	if len(matches) != 6 {
		return instruction{}, fmt.Errorf("invalid instruction input: %q", value)
	}

	var mode modeType
	switch matches[1] {
	case "turn on":
		mode = on
	case "turn off":
		mode = off
	default:
		mode = toggle
	}

	fromX, err := strconv.Atoi(matches[2])
	if err != nil {
		return instruction{}, fmt.Errorf("invalid instruction coordinate: %q", matches[2])
	}
	fromY, err := strconv.Atoi(matches[3])
	if err != nil {
		return instruction{}, fmt.Errorf("invalid instruction coordinate: %q", matches[3])
	}
	toX, err := strconv.Atoi(matches[4])
	if err != nil {
		return instruction{}, fmt.Errorf("invalid instruction coordinate: %q", matches[4])
	}
	toY, err := strconv.Atoi(matches[5])
	if err != nil {
		return instruction{}, fmt.Errorf("invalid instruction coordinate: %q", matches[5])
	}

	if fromX > toX || fromY > toY {
		return instruction{}, fmt.Errorf("invalid instruction input: %q", value)
	}

	return instruction{mode: mode, fromRow: fromX, toRow: toX, fromCol: fromY, toCol: toY}, nil
}
