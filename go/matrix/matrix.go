package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(in string) (Matrix, error) {
	rows := strings.Split(in, "\n")
	matrix := make(Matrix, len(rows))
	for i, row := range rows {
		row = strings.TrimSpace(row)
		items := strings.Split(row, " ")
		if i > 0 {
			if len(items) != len(matrix[i-1]) {
				return nil, errors.New(fmt.Sprintf("Uneven row count"))
			}
		}
		matrix[i] = make([]int, len(items))
		for ii, item := range items {
			num, err := strconv.Atoi(item)
			if err == nil {
				matrix[i][ii] = num
			} else {
				return nil, errors.New(fmt.Sprintf("Unable to convert string to number"))
			}
		}
	}
	return matrix, nil
}

func (trix *Matrix) Set(row, col, val int) (ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()

	ok = true
	(*trix)[row][col] = val
	return ok
}

func (trix Matrix) Rows() [][]int {
	rows := make([][]int, len(trix))
	for i, row := range trix {
		rows[i] = make([]int, len(row))
		for ii, item := range row {
			rows[i][ii] = item
		}
	}

	return rows
}

func (trix Matrix) Cols() [][]int {
	cols := make([][]int, len(trix[0]))
	row_number := len(trix)
	col_number := len(trix[0])

	for i := 0; i < col_number; i++ {
		cols[i] = make([]int, len(trix))
	}

	for i := 0; i < row_number; i++ {
		for ii := 0; ii < col_number; ii++ {
			cols[ii][i] = trix[i][ii]
		}
	}

	return cols
}
