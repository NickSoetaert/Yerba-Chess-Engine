package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

func getNormalKingMoves(kings, ownPieces, attackedSquares uint64, ch chan []Move) {
	var moves []Move
	kings &= ownPieces
	baseMove := Move(0)
	baseMove.setMoveType(normalKingMove)
	currentSquare := uint8(bits.TrailingZeros64(kings))
	baseMove.setOriginFromSquare(currentSquare)
	possibleAttacks := KingMask[currentSquare]
	possibleAttacks = possibleAttacks &^ attackedSquares &^ ownPieces

	for possibleAttacks != 0 {
		move := baseMove
		attack := utils.IsolateLsb(possibleAttacks)
		move.setDestFromBB(attack)
		possibleAttacks ^= attack
		moves = append(moves, move)
	}
	ch <- moves
}

func (b *Board) getCastlingMoves(ch chan []Move) {
	var moves []Move
	if b.IsWhiteMove {
		if b.WhiteKingsideCastleRights {
			//check if any pieces are on f1 or g1
			if (b.Black | b.White) & (F1 | G1) == 0 {
				//Todo: check if e1, f1, or g1 are under attack
				moves = append(moves, Move(castleKingside))
			}
		}
		if b.WhiteQueensideCastleRights {
			//check if any pieces are on d1, c1 or b1
			if (b.Black | b.White) & (D1 | C1 | B1) == 0 {
				//Todo: check if e1, d1, or c1 are under attack
				moves = append(moves, Move(castleQueenside))
			}
		}
	} else {
		if b.BlackKingsideCastleRights {
			//check if e8, f8, or g8 are under attack
			//check if any pieces are on f8 or g8
		}
		if b.WhiteQueensideCastleRights {
			//check if e8, d8, or c8 are under attack
			//check if any pieces are on d8, c8 or b8
		}
	}
	ch <- moves
}