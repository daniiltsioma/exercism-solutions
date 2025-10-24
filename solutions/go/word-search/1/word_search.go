package wordsearch

import (
	"errors"
	"strings"
)

var ErrNotFound = errors.New("Not found")

// reverseString reverses the string.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// findWord finds the coordinates of the word in the puzzle.
func findWord(word string, puzzle []string) ([2][2]int, error) {
	for r, row := range puzzle {
		// look horizontally LTR
		if idx := strings.Index(row, word); idx > -1 {
			return [2][2]int{
				{idx, r},
				{idx+len(word)-1, r},
			}, nil
		}
		// look horizontally RTL
		if idx := strings.Index(row, reverseString(word)); idx > -1 {
			return [2][2]int{
				{idx+len(word)-1, r},
				{idx, r},
			}, nil
		}

		for c, char := range row {
			if char == rune(word[0]) {
				// look vertically TTB
				i := 1
				for i < len(word) && r+i < len(puzzle) {
					if puzzle[r+i][c] != word[i] {
						break
					}
					i++
				}
				if i == len(word) {
					return [2][2]int{
						{c, r},
						{c, r+len(word)-1},
					}, nil
				}

				// look vertically BTT
				i = 1
				for i < len(word) && r-i >= 0 {
					if puzzle[r-i][c] != word[i] {
						break
					}
					i++
				}
				if i == len(word) {
					return [2][2]int{
						{c, r},
						{c, r-len(word)+1},
					}, nil
				}

				// look diagonally :0

				// from top left
				i = 1
				for i < len(word) && r+i < len(puzzle) && c+i < len(row) {
					if puzzle[r+i][c+i] != word[i] {
						break
					}
					i++
				}
				if i == len(word) {
					return [2][2]int{
						{c, r},
						{c+len(word)-1, r+len(word)-1},
					}, nil
				}

				// from bottom left
				i = 1
				for i < len(word) && r-i >= 0 && c+i < len(row) {
					if puzzle[r-i][c+i] != word[i] {
						break
					}
					i++
				}
				if i == len(word) {
					return [2][2]int{
						{c, r},
						{c+len(word)-1, r-len(word)+1},
					}, nil
				}

				// from top right
				i = 1
				for i < len(word) && r+i < len(puzzle) && c-i >= 0 {
					if puzzle[r+i][c-i] != word[i] {
						break
					}
					i++
				}
				if i == len(word) {
					return [2][2]int{
						{c, r},
						{c-len(word)+1, r+len(word)-1},
					}, nil
				}

				// from bottom right
				i = 1
				for i < len(word) && r-i >= 0 && c-i >= 0 {
					if puzzle[r-i][c-i] != word[i] {
						break
					}
					i++
				}
				if i == len(word) {
					return [2][2]int{
						{c, r},
						{c-len(word)+1, r-len(word)+1},
					}, nil
				}
			}
		}
	}

	// word not found
	return [2][2]int{{-1, -1}, {-1,-1}}, ErrNotFound
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	res := map[string][2][2]int{}

	var e error = nil

	for _, word := range words {
		point, err := findWord(word, puzzle)
		if err != nil {
			e = err
		}
		res[word] = point
	}

	return res, e
}
