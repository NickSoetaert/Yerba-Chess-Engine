package utils

import (
	"strconv"
)
/*
GetBoardKey takes a board state for a single piece, and returns
the long representation. For debug upropses only.
*/
func GetBoardKey() uint64 {
	board := [8][8]string{
		{"x","x","x"," "," "," "," "," "},
		{"x"," ","x"," "," "," "," "," "},
		{"x","x","x"," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "},
	}
	board[0][0]="y"

	var result uint64 
	var str string

	for i := 0; i < 64; i++ {
		if board[i/8][i%8] != " " {
			result += 0
		}
	}


	return result
}
