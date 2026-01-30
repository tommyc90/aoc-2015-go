package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tommyc90/aoc-2015-go/internal/solutions"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: app <day> <part> <input-file>")
		return
	}

	day := os.Args[1]
	part := os.Args[2]
	file := os.Args[3]

	if part != "1" && part != "2" {
		fmt.Println("Invalid part. Only part 1 and part 2 are supported.")
		return
	}

	start := time.Now()
	defer func() {
		fmt.Printf("Execution time: %s\n", time.Since(start))
	}()

	solver, ok := solutions.Get(day)
	if !ok {
		fmt.Println("Day not implemented.")
		return
	}

	switch part {
	case "1":
		solver.SolvePart1(file)
	case "2":
		solver.SolvePart2(file)
	}
}
