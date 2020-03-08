package moveGen

func GetPawnMoves(pawns, friends, enemies uint64, isWhiteToMove bool) uint64 {
	return unfilteredPawnNonAttacks(pawns&friends, isWhiteToMove)
}

//TODO: shifting them all together doesn't really work - I need to know which pawns can go where.
func unfilteredPawnNonAttacks(pawns uint64, isWhiteToMove bool) (moves uint64) {
	if isWhiteToMove {
		moves = pawns << 8
		moves |= (pawns << 16) & FourthRank //You can only move 2 squares if you end up on the 4th rank for white.
	} else {
		moves = pawns >> 8
		moves |= (pawns >> 16) & FifthRank
	}
	return moves
}