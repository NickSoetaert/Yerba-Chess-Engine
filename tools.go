package main

import (
	"fmt"
	"math/bits"
)

/*
GetBoardKey takes a board state for a single piece, and returns
the long representation. For debug upropses only.
*/
func GetBoardKey() uint64 {
	b := [8][8]string{
		{" ", " ", " ", " ", " ", " ", " ", " "}, //8
		{" ", " ", " ", " ", " ", " ", " ", " "}, //7
		{" ", " ", " ", " ", " ", " ", " ", " "}, //6
		{" ", " ", " ", " ", " ", "x", " ", " "}, //5
		{" ", " ", " ", " ", "x", " ", " ", " "}, //4
		{" ", " ", " ", " ", " ", " ", " ", " "}, //3
		{" ", " ", " ", " ", "x", " ", " ", " "}, //2
		{" ", " ", " ", " ", " ", "x", " ", " "}, //1
		//A    B    C    D    E    F    G    H
	}

	var result uint64

	for i := uint8(0); i < 64; i++ {
		if b[7-(i/8)][i%8] != " " {
			result += 1 << i
		}
	}

	return result << 1
}

//PrintBinaryBoard takes a bitboard and prints it in chess-board format
func PrintBinaryBoard(b uint64) {
	mask := A8
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

func UPopCount(b uint64) uint {
	return uint(bits.OnesCount64(b))
}

func U8PopCount(b uint64) uint8 {
	return uint8(bits.OnesCount64(b))
}