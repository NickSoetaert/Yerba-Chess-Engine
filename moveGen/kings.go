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

func (b *Board) getCastlingMoves(attackedSquares uint64, ch chan []Move) {
	var moves []Move
	if b.IsWhiteMove {
		if b.WhiteKingsideCastleRights {
			if (b.BlackPieces|b.WhitePieces)&(F1|G1) == 0 && attackedSquares&(E1|F1|G1) == 0 {
				moves = append(moves, Move(castleKingside))
			}
		}
		if b.WhiteQueensideCastleRights {
			if (b.BlackPieces|b.WhitePieces)&(D1|C1|B1) == 0 && attackedSquares&(E1|D1|C1) == 0 {
				moves = append(moves, Move(castleQueenside))
			}
		}
	} else {
		if b.BlackKingsideCastleRights {
			if (b.BlackPieces|b.WhitePieces)&(F8|G8) == 0 && attackedSquares&(E8|F8|G8) == 0 {
				moves = append(moves, Move(castleQueenside))
			}
		}
		if b.WhiteQueensideCastleRights {
			if (b.BlackPieces|b.WhitePieces)&(D8|C8|B8) == 0 && attackedSquares&(E8|D8|C8) == 0 {
				moves = append(moves, Move(castleQueenside))
			}
		}
	}
	ch <- moves
}
