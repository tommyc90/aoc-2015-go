package day09

import (
	"fmt"
	"log"
	"math"

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

	locations, err := parseLocations(lines)
	if err != nil {
		log.Fatal(err)
	}

	d := findShortestDistance(locations)

	fmt.Println(d)
}

func (s *Solver) SolvePart2(filePath string) {
	lines, err := utils.ReadLines(filePath)
	if err != nil {
		log.Fatal(err)
	}

	locations, err := parseLocations(lines)
	if err != nil {
		log.Fatal(err)
	}

	d := findLongestDistance(locations)

	fmt.Println(d)
}

func findShortestDistance(locations map[string]*Location) int {
	distance := math.MaxInt
	toVisit := make(map[string]struct{})
	for lName := range locations {
		toVisit[lName] = struct{}{}
	}

	for _, location := range locations {
		d, ok := findShortestDistanceFrom(location, toVisit, 0)
		if !ok {
			continue
		}
		if d < distance {
			distance = d
		}
	}

	return distance
}

func findShortestDistanceFrom(location *Location, toVisit map[string]struct{}, currentDistance int) (int, bool) {
	_, ok := toVisit[location.Name]
	if !ok {
		return 0, false // already visited
	}

	delete(toVisit, location.Name)
	defer func() {
		toVisit[location.Name] = struct{}{}
	}()

	if len(toVisit) == 0 {
		return currentDistance, true
	}

	traversed := false
	distance := math.MaxInt
	for _, route := range location.Routes {
		d, ok := findShortestDistanceFrom(route.Location, toVisit, currentDistance+route.Distance)
		if !ok {
			continue
		}
		if d < distance {
			distance = d
			traversed = true
		}
	}

	if !traversed {
		return 0, false // dead end
	}

	return distance, true
}

func findLongestDistance(locations map[string]*Location) int {
	distance := math.MinInt
	toVisit := make(map[string]struct{})
	for lName := range locations {
		toVisit[lName] = struct{}{}
	}

	for _, location := range locations {
		d, ok := findLongestDistanceFrom(location, toVisit, 0)
		if !ok {
			continue
		}
		if d > distance {
			distance = d
		}
	}

	return distance
}

func findLongestDistanceFrom(location *Location, toVisit map[string]struct{}, currentDistance int) (int, bool) {
	_, ok := toVisit[location.Name]
	if !ok {
		return 0, false // already visited
	}

	delete(toVisit, location.Name)
	defer func() {
		toVisit[location.Name] = struct{}{}
	}()

	if len(toVisit) == 0 {
		return currentDistance, true
	}

	traversed := false
	distance := math.MinInt
	for _, route := range location.Routes {
		d, ok := findLongestDistanceFrom(route.Location, toVisit, currentDistance+route.Distance)
		if !ok {
			continue
		}
		if d > distance {
			distance = d
			traversed = true
		}
	}

	if !traversed {
		return 0, false // dead end
	}

	return distance, true
}
