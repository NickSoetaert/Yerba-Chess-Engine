package utils

import (
	//"strconv"
	"fmt"
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
		{" "," "," "," ","x"," "," "," "},
		{" "," "," "," "," "," "," "," "},
		{" "," "," "," "," "," "," "," "},
	}
	board[0][0]="y"

	var result uint64 
	//var str string

	for i := uint8(0); i < 64; i++ {
		if board[i/8][i%8] != " " {
			result += 1 << i
		}
	}


	return result
}


func GetSingleBB(key uint64){
	board := [8][8]string{{}}
	str := strconv.FormatUint(key, 2)


	//fmt.Println(len(str))
	//fmt.Println(str[63])

	//for i:= len(str)-1; i >= 0; i-- {
	//bug: str isn't len(64) unless there's pieces on a1 and h8

	for i:= 0; i < len(str); i++{

		if str[i] == '1'{
			board[i/8][i%8] = "x"
		} else {
			board[i/8][i%8] = " "
		}
	}
	
	for _, row := range board{
		for _, cell := range row{
			print("|",cell)
		}
		fmt.Println("|")
	}
}

func Test() {
	x := GetBoardKey()
	GetSingleBB(x)
	fmt.Println("a")
}