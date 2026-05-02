package day08

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolvePart1(filePath string) {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	diff := 0
	for _, line := range lines {
		diff += part1(line)
	}

	fmt.Println(diff)
}

func (s *Solver) SolvePart2(filePath string) {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	diff := 0
	for _, line := range lines {
		diff += part2(line)
	}

	fmt.Println(diff)
}

var simpleEscape = regexp.MustCompile(`\\[\\"]`)
var hexEscape = regexp.MustCompile(`\\x[0-9a-f]{2}`)

func part1(line string) int {
	sm := simpleEscape.FindAllString(line, -1)
	l2 := simpleEscape.ReplaceAllString(line, "_")
	hm := hexEscape.FindAllString(l2, -1)

	return len(sm) + len(hm)*3 + 2
}

func part2(line string) int {
	l1 := line[1 : len(line)-1]
	toEcapeCount := strings.Count(l1, "\"") + strings.Count(l1, "\\")

	return toEcapeCount + 4
}
