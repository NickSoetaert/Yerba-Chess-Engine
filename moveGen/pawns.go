package moveGen

import (
	"Yerba/utils"
)

func GetPawnMoves(pawns, whitePieces, blackPieces uint64, isWhiteToMove bool) (moves []uint64) {
	if isWhiteToMove {
		pawns &= whitePieces
	} else {
		pawns &= blackPieces
	}

	for i := 0; pawns != 0; i++ {
		pawn := utils.IsolateLsb(pawns)
		moves = append(moves, pawnNormalCapturingMoves(pawn, whitePieces, blackPieces, isWhiteToMove)|
			pawnNonAttacks(pawn, whitePieces, blackPieces, isWhiteToMove))

		pawns &= pawns - 1
	}
	return moves
}
//TODO for all move generators - return an array of Moves instead of a bitboard (or array of BBs)
//
func pawnNonAttacks(pawn, whitePieces, blackPieces uint64, isWhiteToMove bool) (moves uint64) {
	if isWhiteToMove {
		moves = pawn << 8
		moves |= (pawn << 16) & FourthRank
	} else {
		moves = pawn >> 8
		moves |= (pawn >> 16) & FifthRank
	}
	return moves & ^(whitePieces | blackPieces)
}

func pawnNormalCapturingMoves(pawn, whitePieces, blackPieces uint64, isWhiteToMove bool) (moves uint64) {
	if isWhiteToMove {
		moves = ((pawn << 7) & ^HFile) & blackPieces
		moves |= ((pawn << 9) & ^AFile) & blackPieces
	} else {
		moves = ((pawn >> 7) & ^AFile) & whitePieces
		moves |= ((pawn >> 9) & ^HFile) & whitePieces
	}
	return
}
