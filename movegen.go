package main

import (
	"math/bits"
)

//KnightAttacks returns a BB of all squares that knights of a given color can attack, regardless of other pieces on the board.
func KnightAttacks(knights BinaryBoard) BinaryBoard {

	var possibleAttacks BinaryBoard
	//Get number of squares that the knight can attack
	count := bits.OnesCount64(uint64(knights))
	//count must be used, else loop range will decrement with each pass.
	for i := 0; i < count; i++ {
		singleKnightPosition := uint8(bits.TrailingZeros64(uint64(knights)))
		possibleAttacks |= knightMoves[singleKnightPosition] //get current square that knight can attack
		knights ^= 1 << singleKnightPosition                 //now clear that knight for the next loop iteration
	}
	return possibleAttacks
}

func RookAttacks(rooks BinaryBoard) BinaryBoard {

	var possibleAttacks BinaryBoard

	count := bits.OnesCount64(uint64(rooks))
	//count must be used, else loop range will decrement with each pass.
	for i := 0; i < count; i++ {
		singleRookPosition := uint8(bits.TrailingZeros64(uint64(rooks)))
		possibleAttacks |= rookMult[singleRookPosition] //get current squares that rook can attack
		rooks ^= 1 << singleRookPosition                //now clear that rook for the next loop iteration
	}

	return possibleAttacks
}

func DownFill(file BinaryBoard) BinaryBoard {
	file |= file >> 8
	file |= file >> 16
	file |= file >> 32
	return file
}