package day02

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type present struct {
	width, height, length int
}

func parsePresent(value string) (present, error) {
	dims, err := parsePresentDims(value)
	if err != nil {
		return present{}, err
	}

	return present{width: dims[0], height: dims[1], length: dims[2]}, nil
}

func (p present) TotalSurfaceArea() int {
	return 2*p.height*p.width + 2*p.height*p.length + 2*p.width*p.length
}

func (p present) MinSideSurfaceArea() int {
	return min(p.height*p.width, p.height*p.length, p.width*p.length)
}

func (p present) MinPerimeter() int {
	dims := [3]int{p.height, p.width, p.length}
	slices.Sort(dims[:])
	return 2 * (dims[0] + dims[1])
}

func (p present) Volume() int {
	return p.height * p.width * p.length
}

func parsePresentDims(value string) ([3]int, error) {
	dims := [3]int{}
	parts := strings.Split(value, "x")
	if len(parts) != 3 {
		return dims, fmt.Errorf("invalid present input: %q", value)
	}

	for i, p := range parts {
		dim, err := strconv.Atoi(p)
		if err != nil {
			return dims, err
		}
		dims[i] = dim
	}

	return dims, nil
}
