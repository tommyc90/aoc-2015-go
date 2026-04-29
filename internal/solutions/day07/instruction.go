package day07

import (
	"fmt"
	"regexp"
	"strconv"
)

type operatorType int

const (
	and operatorType = iota
	or
	lshift
	rshift
	not
	wire
)

type operand string

var operandSignalCheck = regexp.MustCompile(`^[a-z]+$`)

func (o operand) isSignal() bool {
	return operandSignalCheck.MatchString(string(o))
}

func (o operand) value() (uint16, error) {
	v, err := strconv.Atoi(string(o))
	if err != nil {
		return 0, fmt.Errorf("invalid operand value: %s", o)
	}

	return uint16(v), nil
}

type instruction struct {
	operator           operatorType
	operand1, operand2 operand
	outputSignal       string
}

var instructionAndOr = regexp.MustCompile(`^([a-z]+|\d+) (AND|OR) ([a-z]+|\d+) -> ([a-z]+)$`)
var instructionShift = regexp.MustCompile(`^([a-z]+) ([LR]SHIFT) (\d+) -> ([a-z]+)$`)
var instructionNot = regexp.MustCompile(`^NOT ([a-z]+) -> ([a-z]+)$`)
var instructionWire = regexp.MustCompile(`^([a-z]+|\d+) -> ([a-z]+)$`)

func parseInstruction(value string) (instruction, error) {
	var operator operatorType

	matches := instructionAndOr.FindStringSubmatch(value)
	if len(matches) == 5 {
		switch matches[2] {
		case "AND":
			operator = and
		default:
			operator = or
		}
		return instruction{
			operator:     operator,
			operand1:     operand(matches[1]),
			operand2:     operand(matches[3]),
			outputSignal: matches[4],
		}, nil
	}

	matches = instructionShift.FindStringSubmatch(value)
	if len(matches) == 5 {
		switch matches[2] {
		case "LSHIFT":
			operator = lshift
		default:
			operator = rshift
		}
		return instruction{
			operator:     operator,
			operand1:     operand(matches[1]),
			operand2:     operand(matches[3]),
			outputSignal: matches[4],
		}, nil
	}

	matches = instructionNot.FindStringSubmatch(value)
	if len(matches) == 3 {
		return instruction{
			operator:     not,
			operand1:     operand(matches[1]),
			outputSignal: matches[2],
		}, nil
	}

	matches = instructionWire.FindStringSubmatch(value)
	if len(matches) == 3 {
		return instruction{
			operator:     wire,
			operand1:     operand(matches[1]),
			outputSignal: matches[2],
		}, nil
	}

	return instruction{}, fmt.Errorf("invalid instruction input: %q", value)
}
