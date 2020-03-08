package utils

import (
	"math/bits"
)


//GetBoardKey takes a board state for a single piece, and returns
//the long representation. For debug purposes only.

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



func UPopCount(b uint64) uint {
	return uint(bits.OnesCount64(b))
}

func U8PopCount(b uint64) uint8 {
	return uint8(bits.OnesCount64(b))
}
