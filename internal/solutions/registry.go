package solutions

import (
	"fmt"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/day01"
	"github.com/tommyc90/aoc-2015-go/internal/solutions/day02"
	"github.com/tommyc90/aoc-2015-go/internal/solutions/day03"
	"github.com/tommyc90/aoc-2015-go/internal/solutions/day04"
	"github.com/tommyc90/aoc-2015-go/internal/solutions/day05"
	"github.com/tommyc90/aoc-2015-go/internal/solutions/day06"
)

var registry = make(map[string]Solver)

func init() {
	register("1", day01.New())
	register("2", day02.New())
	register("3", day03.New())
	register("4", day04.New())
	register("5", day05.New())
	register("6", day06.New())
}

func register(day string, solver Solver) {
	if _, exists := registry[day]; exists {
		panic(fmt.Sprintf("Solver for day %s already registered", day))
	}
	registry[day] = solver
}

func Get(day string) (Solver, bool) {
	solver, ok := registry[day]
	return solver, ok
}
