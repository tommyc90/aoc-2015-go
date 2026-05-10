package day09

import "testing"

type WantRoutes struct {
	from     string
	to       string
	distance int
}

func TestParse(t *testing.T) {
	tests := []struct {
		input string
		want  []WantRoutes
	}{
		{
			input: `Alpha to Beta = 123`,
			want: []WantRoutes{
				{
					from:     "Alpha",
					to:       "Beta",
					distance: 123,
				},
				{
					from:     "Beta",
					to:       "Alpha",
					distance: 123,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			locations, err := parseLocations([]string{tt.input})
			if err != nil {
				t.Fatal("unexpected error")
			}

			for _, w := range tt.want {
				l, ok := locations[w.from]
				if !ok {
					t.Errorf("location %q not found", w.from)
				}
				r, ok := l.Routes[w.to]
				if !ok {
					t.Errorf("route %q not found", w.to)
				}
				if r.Distance != w.distance {
					t.Errorf("unexpected distance: got %d, want %d", r.Distance, w.distance)
				}
			}
		})
	}
}

func TestParse_InvalidInput(t *testing.T) {
	tests := []struct {
		input string
	}{
		{input: `Alpha to Beta != 123`},
		{input: `Alpha to Beta != -123`},
		{input: `1 to 2 = 3`},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			_, err := parseLocations([]string{tt.input})
			if err == nil {
				t.Fatal("expected error, got nil")
			}
		})
	}
}
