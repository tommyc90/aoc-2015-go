package day06

import (
	"fmt"
	"log"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolvePart1(filePath string) {
	instructions, err := readInstructions(filePath)
	if err != nil {
		log.Fatal(err)
	}

	g := newGrid(1000, 1000)
	for _, i := range instructions {
		g.applyInstruction(i)
	}

	fmt.Println(g.countLights())
}

func (s *Solver) SolvePart2(filePath string) {
	instructions, err := readInstructions(filePath)
	if err != nil {
		log.Fatal(err)
	}

	g := newGrid2(1000, 1000)
	for _, i := range instructions {
		g.applyInstruction(i)
	}

	fmt.Println(g.brightness())
}

func readInstructions(filePath string) ([]instruction, error) {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		return nil, err
	}

	instructions := make([]instruction, len(lines))
	for i, l := range lines {
		ins, err := parseInstruction(l)
		if err != nil {
			return nil, err
		}
		instructions[i] = ins
	}

	return instructions, nil
}
