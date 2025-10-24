package matrix

import (
	"errors"
	"strings"
	"strconv"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (*Matrix, error) {
	var m Matrix

	if len(s) == 0 {
		return &Matrix{}, nil
	}

	// split into rows.
	rows := strings.Split(s, "\n")
	m = make([][]int, len(rows))
	// determine the number of columns from the first row.
	// trim spaces from the row beforehand.
	firstRow := strings.Trim(rows[0], " ")
	columns := len(strings.Split(firstRow, " "))
	// iterate through rows.
	for r, row := range rows {
		// trim spaces from the row.
		row = strings.Trim(row, " ")
		// check that row is not empty.
		if len(row) == 0 {
			return nil, errors.New("empty row")
		}
		// validate the number of columns.
		rowColumns := strings.Split(row, " "); if len(rowColumns) != columns {
			return nil, errors.New("uneven rows")
		}
		// allocate columns.
		m[r] = make([]int, columns)
		// iterate through columns.
		for c, column := range rowColumns {
			n, err := strconv.ParseInt(column, 10, 64)
			if err != nil {
				return nil, errors.New("error parsing the matrix")
			}
			m[r][c] = int(n)
		}	
	}

	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	// get the number of rows and columns from the matrix.
	matrixRows := len(m)
	matrixColumns := len(m[0])

	columns := make([][]int, matrixColumns) 
	for i := 0; i < matrixColumns; i++ {
		columns[i] = make([]int, matrixRows)
	}

	for r, row := range m {
		for c, val := range row {
			columns[c][r] = int(val)
		}
	}

	return columns
}

func (m Matrix) Rows() [][]int {
	rows := [][]int{}

	for _, row := range m {
		r := []int{}
		for _, val := range row {
			r = append(r, val)
		}	
		rows = append(rows, r)
	}
	return rows
}

func (m Matrix) Set(row, col, val int) bool {
	// check for invalid row.
	if row < 0 || len(m) - 1 < row {
		return false
	}
	// check for invalid column.
	if col < 0 || len(m[row]) - 1 < col {
		return false
	}
	m[row][col] = val
	return true
}
