package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

func getNormalKingMoves(kings, ownPieces, attackedSquares uint64, ch chan []Move) {
	var moves []Move
	kings &= ownPieces
	baseMove := Move(0)
	baseMove.setMoveType(normalMove)
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
	allPieces := b.BlackPieces & b.WhitePieces

	if b.IsWhiteMove && b.WhiteKingHasNeverMoved {
		if b.H1RookHasNeverMoved {	//Try to castle kingside

			//If there are no blocking pieces, and the king is not in or will be traveling through check
			if allPieces & (F1|G1) == 0 && allPieces & (E1|F1|G1) == 0 {
				var move Move
				move.setOriginFromBB(E1) //all that's needed to update castling rights
				move.setMoveType(castleKingside)
				moves = append(moves, move)
			}
		}

		if b.H1RookHasNeverMoved{	//Try to castle queenside
			if allPieces & (D1|C1|B1) == 0 && attackedSquares&(E1|D1|C1) == 0 {
				var move Move
				move.setOriginFromBB(E1) //all that's needed to update castling rights
				move.setMoveType(castleQueenside)
				moves = append(moves, move)
			}
		}
	}

	//If black's move and their king hasn't moved yet
	if (! b.IsWhiteMove) && b.BlackKingHasNeverMoved {
		if b.H8RookHasNeverMoved { //try to castle kingside
			if (b.BlackPieces|b.WhitePieces)&(F8|G8) == 0 && attackedSquares&(E8|F8|G8) == 0 {
				var move Move
				move.setOriginFromBB(E8) //all that's needed to update castling rights
				move.setMoveType(castleKingside)
				moves = append(moves, move)
			}
		}
		if b.A8RookHasNeverMoved {
			if (b.BlackPieces|b.WhitePieces)&(D8|C8|B8) == 0 && attackedSquares&(E8|D8|C8) == 0 {
				var move Move
				move.setOriginFromBB(E8) //all that's needed to update castling rights
				move.setMoveType(castleQueenside)
				moves = append(moves, move)
			}
		}
	}
	ch <- moves
}
