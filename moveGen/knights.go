package moveGen

import "math/bits"

//UnfilteredKnightAttacks returns a BB of all squares that knights of a given color can attack, regardless of other pieces on the board.
func UnfilteredKnightAttacks(knights uint64) uint64 {
	var possibleAttacks uint64
	//Get number of squares that the knight can attack
	count := bits.OnesCount64(knights)
	for i := 0; i < count; i++ {
		currentSquare := uint8(bits.TrailingZeros64(knights))
		possibleAttacks |= KnightMask[currentSquare] //get current square that knight can attack
		knights ^= 1 << currentSquare                //now clear that knight for the next loop iteration
	}
	return possibleAttacks
}

