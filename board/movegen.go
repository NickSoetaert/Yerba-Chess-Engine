package board

import (
	"math/bits"
)

//KnightAttacks returns a BB of all squares that knights of a given color can attack, regardless of other pieces on the board.
func KnightAttacks(knights BinaryBoard) BinaryBoard {

	var attackers BinaryBoard
	count := bits.OnesCount64(uint64(knights))
	//count must be used, else loop range will decrement with each pass.
	for i := 0; i < count; i++ {
		singleKnightPosition := uint8(bits.TrailingZeros64(uint64(knights)))
		attackers |= knightMoves[singleKnightPosition] //get all squares that knight can attack
		knights ^= (1 << singleKnightPosition)         //now clear that knight for the next loop iteration
	}
	return attackers
}
