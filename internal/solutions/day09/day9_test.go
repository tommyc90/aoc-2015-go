package day09

import "testing"

func TestPart1(t *testing.T) {
	locations, err := parseLocations([]string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	})
	want := 605

	if err != nil {
		t.Fatal("unexpected error")
	}

	d := findShortestDistance(locations)
	if d != want {
		t.Errorf("got %d, want %d", d, want)
	}
}

func TestPart2(t *testing.T) {
	locations, err := parseLocations([]string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	})
	want := 982

	if err != nil {
		t.Fatal("unexpected error")
	}

	d := findLongestDistance(locations)
	if d != want {
		t.Errorf("got %d, want %d", d, want)
	}
}
