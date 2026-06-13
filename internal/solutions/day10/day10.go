package day10

import (
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

	length := runConway(input, 40)

	fmt.Println(length)
}

func (s *Solver) SolvePart2(filePath string) {
	input, err := utils.ReadSingleLine(filePath)
	if err != nil {
		log.Fatal(err)
	}

	length := runConway(input, 50)

	fmt.Println(length)
}

func runWithCache(input string, times int) int {
	const SPLIT_THRESHOLD = 1000
	cache := make(map[string]string)
	for range times {
		parts := split(input, SPLIT_THRESHOLD)
		input = ""
		for _, p := range parts {
			next, ok := cache[p]
			if !ok {
				next = progress(p)
				cache[p] = next
			}
			input += next
		}
	}

	return len(input)
}

func runConway(input string, times int) int {
	elements := make(map[string]int)
	elements[input] = 1
	for range times {
		nextElements := make(map[string]int)
		for e, c := range elements {
			des, err := Decay(e)
			if err != nil {
				log.Fatal(err)
			}
			for _, de := range des {
				_, ok := nextElements[de]
				if ok {
					nextElements[de] += c
					continue
				}
				nextElements[de] = c
			}
		}
		elements = nextElements
	}

	totalLength := 0
	for e, c := range elements {
		totalLength += len(e) * c
	}
	return totalLength
}

func progress(input string) string {
	output := ""
	length := len(input)
	var count, i int
	var current byte
	for {
		var c byte
		if i < length {
			c = input[i]
		}

		if i != 0 && current != c {
			output += strconv.Itoa(count) + string(current)
			count = 0
		}

		current = c
		count++

		if i == length {
			break
		}

		i++
	}

	return output
}

func split(input string, threshold int) []string {
	parts := make([]string, 0)
	length := len(input)
	var part string
	var next rune
	for i, c := range input {
		part += string(c)
		next = 0
		if i+1 < length {
			next = rune(input[i+1])
		}
		if i+1 == length || (c == '3' && next != '3' && len(part) >= threshold) {
			parts = append(parts, part)
			part = ""
		}
	}

	return parts
}
