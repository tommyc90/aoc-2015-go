package solutions

import (
	"fmt"

	"github.com/tommyc90/aoc-2015-go/internal/solutions/day01"
)

var registry = make(map[string]Solver)

func init() {
	register("1", day01.New())
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
