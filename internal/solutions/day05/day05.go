package day05

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

	niceCount := 0
	for _, line := range lines {
		if part1(line) {
			niceCount++
		}
	}

	fmt.Println(niceCount)
}

func (s *Solver) SolvePart2(filePath string) {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	niceCount := 0
	for _, line := range lines {
		if part2(line) {
			niceCount++
		}
	}

	fmt.Println(niceCount)
}

var forbiddenRe = regexp.MustCompile("(ab|cd|pq|xy)")

func part1(input string) bool {
	if forbiddenRe.MatchString(input) {
		return false
	}

	vowelCount := 0
	hasRepeatedLetter := false
	for i := 0; i < len(input); i++ {
		c := input[i]

		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++
		}

		if !hasRepeatedLetter && i > 0 {
			hasRepeatedLetter = c == input[i-1]
		}
	}

	return vowelCount >= 3 && hasRepeatedLetter
}

func part2(input string) bool {
	hasRepeatedLetter := false
	hasRepeatedPair := false
	for i := 0; i < len(input); i++ {
		c := input[i]

		if !hasRepeatedLetter && i > 1 {
			hasRepeatedLetter = c == input[i-2]
		}

		if !hasRepeatedPair && i > 0 {
			pair := input[i-1 : i+1]
			li := strings.LastIndex(input, pair)
			hasRepeatedPair = li > i
		}
	}

	return hasRepeatedLetter && hasRepeatedPair
}
