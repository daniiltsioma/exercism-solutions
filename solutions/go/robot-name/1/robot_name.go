package robotname

import (
	"math/rand"
	"errors"
)

type Robot struct {
	name string
}

var robotNames = map[string]bool{}

func (r *Robot) Name() (string, error) {	
	// check if robot already has a name.
	if r.name != "" {
		return r.name, nil
	}

	// check if all 5-rune names are already allocated.
	if len(robotNames) == 26 * 26 * 1000 {
		return "", errors.New("All names already allocated.")
	}

	runes := []rune{}
	
	for len(runes) == 0 || robotNames[string(runes)] {
		runes = runes[:0]
		for i := 0; i < 2; i++ {
			runes = append(runes, rune(65 + rand.Intn(26)))
		}
		for i := 0; i < 3; i++ {
			runes = append(runes, rune(48 + rand.Intn(10)))
		}
	}

	newName := string(runes)
	robotNames[newName] = true
	r.name = newName
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
	delete(robotNames, r.name)

	r.Name()
}
