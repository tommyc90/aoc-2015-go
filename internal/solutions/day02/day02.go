package day02

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
	result := part1(readPresents(filePath))

	fmt.Println(result)
}

func (s *Solver) SolvePart2(filePath string) {
	result := part2(readPresents(filePath))

	fmt.Println(result)
}

func readPresents(filePath string) []present {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	presents := make([]present, len(lines))
	for i, line := range lines {
		p, err := parsePresent(line)
		if err != nil {
			log.Fatal(err)
		}
		presents[i] = p
	}

	return presents
}

// part1 computes total wrapping paper needed
func part1(presents []present) int {
	total := 0
	for _, p := range presents {
		total += p.TotalSurfaceArea() + p.MinSideSurfaceArea()
	}

	return total
}

// part2 computes total ribbon length needed
func part2(presents []present) int {
	total := 0
	for _, p := range presents {
		total += p.MinPerimeter() + p.Volume()
	}

	return total
}
