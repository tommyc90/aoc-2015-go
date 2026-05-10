package day09

import (
	"fmt"
	"regexp"
	"strconv"
)

type Location struct {
	Name   string
	Routes map[string]Route
}

type Route struct {
	Location *Location
	Distance int
}

var routeRe = regexp.MustCompile(`^([A-Za-z]+) to ([A-Za-z]+) = (\d+)$`)

func parseLocations(routesInput []string) (map[string]*Location, error) {
	locations := make(map[string]*Location)

	for _, routeInput := range routesInput {
		matches := routeRe.FindStringSubmatch(routeInput)
		if len(matches) != 4 {
			return nil, fmt.Errorf("invalid route input: %q", routeInput)
		}

		distance, err := strconv.Atoi(matches[3])
		if err != nil {
			return nil, fmt.Errorf("invalid route distance: %q", routeInput)
		}

		l1 := addLocation(locations, matches[1])
		l2 := addLocation(locations, matches[2])

		addRoute(l1, l2, distance)
		addRoute(l2, l1, distance)
	}

	return locations, nil
}

func addLocation(locations map[string]*Location, locationName string) *Location {
	location, ok := locations[locationName]
	if ok {
		return location
	}

	location = &Location{
		Name:   locationName,
		Routes: make(map[string]Route),
	}

	locations[locationName] = location

	return location
}

func addRoute(from *Location, to *Location, distance int) {
	if from == nil || to == nil {
		return
	}

	route, ok := from.Routes[to.Name]
	if ok {
		return
	}

	route = Route{
		Location: to,
		Distance: distance,
	}

	from.Routes[to.Name] = route
}
