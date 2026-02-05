package day04

import (
	"crypto/md5"
	"fmt"
	"log"
	"strconv"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/utils"
)

type Solver struct{}

func New() *Solver {
	return &Solver{}
}

func (s *Solver) SolvePart1(filePath string) {
	input, err := utils.ReadSingleLine(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result := part1(input)

	fmt.Println(result)
}

func (s *Solver) SolvePart2(filePath string) {
	input, err := utils.ReadSingleLine(filePath)
	if err != nil {
		log.Fatal(err)
	}

	result := part2(input)

	fmt.Println(result)
}

func part1(input string) int64 {
	return findSuffix([]byte(input), func(s [16]byte) bool {
		// first 20 bits are zero (5 hex zeroes)
		return s[0] == 0 && s[1] == 0 && s[2]&0xF0 == 0
	})
}

func part2(input string) int64 {
	return findSuffix([]byte(input), func(s [16]byte) bool {
		// first 24 bits are zero (6 hex zeroes)
		return s[0] == 0 && s[1] == 0 && s[2] == 0
	})
}

func findSuffix(base []byte, matches func([16]byte) bool) int64 {
	bl := len(base)
	buf := make([]byte, bl+20)
	copy(buf, base)
	for suffix := int64(1); ; suffix++ {
		b := buf[:bl]
		d := strconv.AppendInt(b, suffix, 10)
		s := md5.Sum(d)
		if matches(s) {
			return suffix
		}
	}
}
