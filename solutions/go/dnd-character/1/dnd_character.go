package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor((float64(score) - 10.0) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	throws := []int{0,0,0,0}
	for i := 0; i < 4; i++ {
		throws[i] = rand.Intn(6) + 1
	}
	sort.Ints(throws)
	return throws[1] + throws[2] + throws[3]
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	constitutionScore := Ability()

	characterScores := Character{
		Strength: Ability(),
		Dexterity: Ability(),
		Constitution: constitutionScore,
		Intelligence: Ability(),
		Wisdom: Ability(),
		Charisma: Ability(),
		Hitpoints: 10 + Modifier(constitutionScore),
	}

	return characterScores
}
