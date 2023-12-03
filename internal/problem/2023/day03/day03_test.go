package day03

import "testing"

func TestIsAdjacent(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
		x     int
		y     int
		is    bool
	}{
		{
			name:  "x",
			input: []string{"5.", ".*"},
			x:     0,
			y:     0,
			is:    true,
		},
		{
			name:  "no",
			input: []string{"...", ".5.", "..."},
			x:     1,
			y:     1,
			is:    false,
		},
		{
			name:  "x,y+1",
			input: []string{"...", ".5.", ".*."},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x-1,y",
			input: []string{"...", "*5.", "..."},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x+1,y",
			input: []string{"...", ".5*", "..."},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x,y-1",
			input: []string{".*.", ".5.", "..."},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x+1,y+1",
			input: []string{"...", ".5.", "..*"},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x-1,y+1",
			input: []string{"...", ".5.", "*.."},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x+1,y-1",
			input: []string{"..*", ".5.", "..."},
			x:     1,
			y:     1,
			is:    true,
		},
		{
			name:  "x-1,y-1",
			input: []string{"*..", ".5.", "..."},
			x:     1,
			y:     1,
			is:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, _, _, _ := IsAdjacent(tc.input, tc.x, tc.y)
			if r != tc.is {
				t.Fail()
			}
		})
	}
}
