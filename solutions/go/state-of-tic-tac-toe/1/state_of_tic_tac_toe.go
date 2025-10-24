package stateoftictactoe

import "errors"

type State string

var ErrInvalid = errors.New("invalid board")

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
)

func StateOfTicTacToe(board []string) (State, error) {
	// top, middle, bottom
	t, m, b := board[0], board[1], board[2]

	cells := make([][]rune, 3)

	win := false
	isFull := true

	xCount := 0
	oCount := 0

	// check diagonals
	if t[0] == m[1] && m[1] == b[2] && t[0] != ' ' {
		win = true
	} else if t[2] == m[1] && m[1] == b[0] && t[2] != ' ' {
		if win {
			return "", ErrInvalid
		}
		win = true
	}

	// check verticals
	for i := range 3 {
		if t[i] == m[i] && m[i] == b[i] && t[i] != ' ' {
			if win {
				return "", ErrInvalid
			}
			win = true
		}
	}

	for i, s := range board {
		cells[i] = make([]rune, 3)
		if s == "XXX" {
			win = true
		}
		if s == "OOO" {
			if win {
				return "", ErrInvalid
			}
			win = true
		}
		for j, r := range s {
			if r == 'X' {
				xCount++
			} else if r == 'O' {
				oCount++
			}

			if r == ' ' {
				isFull = false
			}
			cells[i][j] = r
		}
	}

	if xCount < oCount || xCount - oCount > 1 {
		return "", ErrInvalid
	}

	if win {
		return Win, nil
	} 

	if isFull {
		return Draw, nil
	}

	return Ongoing, nil
}