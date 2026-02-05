package day01

import (
	"fmt"
	"log"
	"strings"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolvePart1(filePath string) {
	dirs, err := utils.ReadSingleLine(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result, err := part1(dirs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func (s *Solver) SolvePart2(filePath string) {
	dirs, err := utils.ReadSingleLine(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result, err := part2(dirs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}

func part1(directions string) (int, error) {
	up := strings.Count(directions, "(")
	down := strings.Count(directions, ")")
	if up+down != len(directions) {
		return 0, fmt.Errorf("invalid input: contains unexpected character")
	}

	return up - down, nil
}

func part2(directions string) (int, error) {
	floor := 0
	pos := 1
	for _, c := range directions {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
			if floor < 0 {
				return pos, nil
			}
		default:
			return 0, fmt.Errorf("invalid input: contains unexpected character")
		}
		pos++
	}

	return 0, fmt.Errorf("never reached basement for input: %s", directions)
}
