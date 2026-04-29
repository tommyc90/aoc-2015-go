package day07

import (
	"fmt"
)

type calculator struct {
	instructions map[string]instruction
	signalValues map[string]uint16
}

func newCalculator(instructions map[string]instruction) *calculator {
	return &calculator{
		instructions: instructions,
		signalValues: make(map[string]uint16, len(instructions)),
	}
}

func (c *calculator) calculate(signal string) (uint16, error) {
	v, ok := c.signalValues[signal]
	if ok {
		return v, nil
	}

	ins, ok := c.instructions[signal]
	if !ok {
		return 0, fmt.Errorf("instruction not found for signal %s", signal)
	}

	if ins.outputSignal != signal {
		return 0, fmt.Errorf("invalid instruction mapping for signal %s", signal)
	}

	var operand1, operand2, result uint16
	var err error

	if ins.operand1 == "" {
		return 0, fmt.Errorf("missing operand 1 in instruction %v", ins)
	}

	if ins.operand1.isSignal() {
		operand1, err = c.calculate(string(ins.operand1))
		if err != nil {
			return 0, nil
		}
	} else {
		operand1, err = ins.operand1.value()
		if err != nil {
			return 0, nil
		}
	}

	if ins.operand2.isSignal() {
		operand2, err = c.calculate(string(ins.operand2))
		if err != nil {
			return 0, nil
		}
	} else if ins.operand2 != "" {
		operand2, err = ins.operand2.value()
		if err != nil {
			return 0, nil
		}
	}

	switch ins.operator {
	case and:
		result = operand1 & operand2
	case or:
		result = operand1 | operand2
	case lshift:
		result = operand1 << operand2
	case rshift:
		result = operand1 >> operand2
	case not:
		result = operand1 ^ 0xffff
	case wire:
		result = operand1
	}

	c.signalValues[signal] = result
	return result, nil
}
