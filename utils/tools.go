package utils

import (
	"Yerba/board"
	"fmt"
	"math/bits"
)

/*
GetBoardKey takes a board state for a single piece, and returns
the long representation. For debug upropses only.
*/
func GetBoardKey() uint64 {
	board := [8][8]string{
		{" ", " ", " ", " ", " ", " ", " ", " "}, //8
		{" ", " ", " ", " ", " ", " ", " ", " "}, //7
		{" ", " ", " ", " ", " ", " ", " ", " "}, //6
		{" ", " ", " ", " ", " ", " ", " ", " "}, //5
		{" ", " ", " ", " ", " ", " ", " ", " "}, //4
		{" ", "x", " ", " ", " ", " ", " ", " "}, //3
		{" ", " ", "x", " ", " ", " ", " ", " "}, //2
		{" ", " ", " ", " ", " ", " ", " ", " "}, //1
		//A    B    C    D    E    F    G    H
	}
	//board[0][0] = "y"

	var result uint64
	//var str string

	for i := uint8(0); i < 64; i++ {
		if board[7-(i/8)][i%8] != " " {
			result += 1 << i
		}
	}

	return result
}

//PrintBinaryBoard takes a bitboard and prints it in chess-board format
func PrintBinaryBoard(b board.BinaryBoard) {
	mask := board.A8
	fmt.Println("  ---------------------------------")
	for i := 8; i >= 1; i-- {
		fmt.Printf("%d |", i)
		for j := 1; j <= 8; j++ {
			if bits.OnesCount64(uint64(b&mask)) == 1 {
				fmt.Print(" X |")
			} else {
				fmt.Print("   |")
			}
			if j != 8 {
				mask = mask << 1
			}
		}
		mask = mask >> 15
		fmt.Println("")
		fmt.Println("  ---------------------------------")
	}
	//print numbers at bottom with 3 spaces of padding
	for i := 'A' - 3; i <= 'H'; i++ {
		if i >= 'A' {
			fmt.Printf(" %c |", i)
		} else if i == 'A'-1 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}
