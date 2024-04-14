package main

import (
	"strconv"
)

func in(data int, v []int) bool {
	for _, value := range v {
		if data == value {
			return true
		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int][]int)
	cols := make(map[int][]int)
	squares := make(map[string][]int)

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			key := strconv.Itoa(i/3) + "," + strconv.Itoa(j/3)

			if in(int(board[i][j]-48), cols[j]) ||
				in(int(board[i][j]-48), rows[i]) ||
				in(int(board[i][j]-48), squares[key]) {
                return false
			}

			rows[i] = append(rows[i], int(board[i][j]-48))
			cols[j] = append(cols[j], int(board[i][j]-48))
			squares[key] = append(squares[key], int(board[i][j]-48))
		}
	}

	return true
}

func main() {

	arg := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	println(isValidSudoku(arg))
}
