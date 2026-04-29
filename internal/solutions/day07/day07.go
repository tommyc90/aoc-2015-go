package day07

import (
	"fmt"
	"log"
	"strconv"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/utils"
)

const targetSignal1 = "a"
const targetSignal2 = "b"

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolvePart1(filePath string) {
	instructions, err := readInstructions(filePath)
	if err != nil {
		log.Fatal(err)
	}

	c := newCalculator(instructions)
	targetValue, err := c.calculate(targetSignal1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(targetValue)
}

func (s *Solver) SolvePart2(filePath string) {
	instructions, err := readInstructions(filePath)
	if err != nil {
		log.Fatal(err)
	}

	c1 := newCalculator(instructions)
	targetValue1, err := c1.calculate(targetSignal1)

	if err != nil {
		log.Fatal(err)
	}

	instructions[targetSignal2] = instruction{
		operator:     wire,
		operand1:     operand(strconv.Itoa(int(targetValue1))),
		outputSignal: targetSignal2,
	}
	c2 := newCalculator(instructions)
	targetValue2, err := c2.calculate(targetSignal1)

	fmt.Println(targetValue2)
}

func readInstructions(filePath string) (map[string]instruction, error) {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	instructions := make(map[string]instruction, len(lines))
	for _, l := range lines {
		ins, err := parseInstruction(l)
		if err != nil {
			return nil, err
		}
		instructions[ins.outputSignal] = ins
	}

	_, hasTarget1 := instructions[targetSignal1]
	if !hasTarget1 {
		return nil, fmt.Errorf("missing target signal 1: %s", targetSignal1)
	}

	_, hasTarget2 := instructions[targetSignal2]
	if !hasTarget2 {
		return nil, fmt.Errorf("missing target signal 2: %s", targetSignal2)
	}

	return instructions, nil
}
