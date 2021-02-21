package main

import "fmt"

func main() {
	//matrix := [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}
	matrix := [][]byte{{'0', '1'}, {'1', '0'}}
	max := maximalSquare(matrix)
	fmt.Println(max)
}

func maximalSquare(matrix [][]byte) int {
	max := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if square := pointMaxSquare(matrix, i, j); square > max {
				max = square
			}
		}
	}
	return max
}

func pointMaxSquare(matrix [][]byte, row, col int) int {
	if matrix[row][col] == 0 {
		return 0
	}
	var count, add int
	for add = 0; row+add < len(matrix) && col+add < len(matrix[0]); add++ {
		count = 0
		for i, j := row+add, col; j < len(matrix[0]) && j <= col+add; j++ {
			count += int(matrix[i][j] - '0')
		}
		for i, j := row, col+add; i < len(matrix) && i < row+add; i++ {
			count += int(matrix[i][j] - '0')
		}
		if count != (add+1)*(add+1)-add*add {
			return add * add
		}
	}
	if count != (add+1)*(add+1)-add*add {
		return add * add
	}
	return 0
}
