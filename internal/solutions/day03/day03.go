package day03

import (
	"fmt"
	"log"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/common"
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
	visited := make(map[common.Point]struct{})
	visited[common.Point{}] = struct{}{}
	p := common.Point{}
	for _, d := range directions {
		err := step(d, &p)
		if err != nil {
			return 0, err
		}

		visited[p] = struct{}{}
	}

	return len(visited), nil
}

func part2(directions string) (int, error) {
	visited := make(map[common.Point]struct{})
	visited[common.Point{}] = struct{}{}
	santa := common.Point{}
	robo := common.Point{}
	for i, d := range directions {
		var p *common.Point
		if i%2 == 0 {
			p = &santa
		} else {
			p = &robo
		}

		err := step(d, p)
		if err != nil {
			return 0, err
		}

		visited[*p] = struct{}{}
	}

	return len(visited), nil
}

func step(direction rune, point *common.Point) error {
	switch direction {
	case '>':
		point.X++
	case '<':
		point.X--
	case '^':
		point.Y++
	case 'v':
		point.Y--
	default:
		return fmt.Errorf("invalid direction: %q", direction)
	}

	return nil
}
