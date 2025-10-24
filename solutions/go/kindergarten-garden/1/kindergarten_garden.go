package kindergarten

import (
	"strings"
	"sort"	
	"errors"
)

// Define the Garden type here.
type Garden struct {
	children map[string][]string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	// copy children and sort in alphabetized order.
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	// map letters to plants.
	plants := map[rune]string{
		'G': "grass",
		'C': "clover",
		'R': "radishes",
		'V': "violets",
	}

	// split diagram into rows.
	rows := strings.Split(diagram, "\n")

	// confirm the right number of rows in diagram.
	if len(rows) != 3 {
		return nil, errors.New("wrong diagram format")
	}
	rows = rows[1:]

	// confirm that rows are of the same length.
	if len(rows[0]) != len(rows[1]) {
		return nil, errors.New("mismatched rows")
	}

	// confirm that rows have even number of cups.
	if len(rows[0]) % 2 != 0 {
		return nil, errors.New("odd number of cups")
	}

	garden := Garden{
		children: map[string][]string{},
	}

	for i, child := range sortedChildren {
		// check for duplicate names.
		_, exists := garden.children[child]; if exists {
			return nil, errors.New("duplicate name")
		}
		// iterate over offsets for the current child.
		for _, p := range [][]int{
			{0, i*2},
			{0, i*2+1},
			{1, i*2},
			{1, i*2+1},
		} {
			y, x := p[0], p[1]
			// check for valid cup codes.
			v, ok := plants[rune(rows[y][x])]; if !ok {
				return nil, errors.New("invalid cup code")
			}
			garden.children[child] = append(garden.children[child], v)
		}
	}

	return &garden, nil
}

// Plants returns plants for a specific child.
func (g *Garden) Plants(child string) ([]string, bool) {
	p, ok := g.children[child]; if !ok {
		return nil, false
	}
	return p, true
}