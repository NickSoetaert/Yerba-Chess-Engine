package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

func (b *Board) getNormalKingMoves(attackedSquares uint64, ch chan []Move) {
	var moves []Move
	var currentSquare uint64
	var possibleAttacks uint64

	baseMove := Move(0)
	baseMove.setMoveType(normalMove)

	if b.IsWhiteMove {
		currentSquare = utils.IsolateLsb(b.Kings & b.WhitePieces)
		possibleAttacks = KingMask[bits.TrailingZeros64(currentSquare)] &^ attackedSquares &^ b.WhitePieces
		baseMove.setOriginOccupancy(whiteKing)
		baseMove.setDestOccupancyAfterMove(whiteKing)
	} else {
		currentSquare = utils.IsolateLsb(b.Kings & b.BlackPieces)
		possibleAttacks = KingMask[bits.TrailingZeros64(currentSquare)] &^ attackedSquares &^ b.WhitePieces
		baseMove.setOriginOccupancy(blackKing)
		baseMove.setDestOccupancyAfterMove(blackKing)
	}
	baseMove.setOriginFromBB(currentSquare)

	for possibleAttacks != 0 {
		move := baseMove
		attack := utils.IsolateLsb(possibleAttacks)

		move.setDestOccupancyBeforeMove(b.getTileOccupancy(attack)) //note the piece (or lack of) that's on the square before we capture
		move.setDestFromBB(attack)

		possibleAttacks ^= attack
		moves = append(moves, move)
	}
	ch <- moves
}

func (b *Board) getCastlingMoves(attackedSquares uint64, ch chan []Move) {
	var moves []Move
	allPieces := b.BlackPieces | b.WhitePieces

	if b.IsWhiteMove && b.WhiteKingHasNeverMoved {
		if b.H1RookHasNeverMoved && (b.WhitePieces&b.Rooks&H1 != 0) { //Try to castle kingside

			//If there are no blocking pieces, and the king is not in or will be traveling through check
			if allPieces&(F1|G1) == 0 && attackedSquares&(E1|F1|G1) == 0 {
				var move Move
				move.setOriginFromBB(E1)
				move.setOriginOccupancy(whiteKing)
				move.setDestOccupancyBeforeMove(empty)
				move.setDestOccupancyAfterMove(whiteKing)
				move.setMoveType(castleKingside)
				moves = append(moves, move)
			}
		}

		if b.H1RookHasNeverMoved && (b.WhitePieces&b.Rooks&A1 != 0) { //Try to castle queenside
			if allPieces&(D1|C1|B1) == 0 && attackedSquares&(E1|D1|C1) == 0 {
				var move Move
				move.setOriginFromBB(E1)
				move.setOriginOccupancy(whiteKing)
				move.setDestOccupancyBeforeMove(empty)
				move.setDestOccupancyAfterMove(whiteKing)
				move.setMoveType(castleQueenside)
				moves = append(moves, move)
			}
		}
	}

	//If black's move and their king hasn't moved yet
	if (!b.IsWhiteMove) && b.BlackKingHasNeverMoved {
		if b.H8RookHasNeverMoved && (b.BlackPieces&b.Rooks&H8 != 0) { //try to castle kingside
			if (b.BlackPieces|b.WhitePieces)&(F8|G8) == 0 && attackedSquares&(E8|F8|G8) == 0 {
				var move Move
				move.setOriginFromBB(E8) //all that's needed to update castling rights
				move.setOriginOccupancy(blackKing)
				move.setDestOccupancyBeforeMove(empty)
				move.setDestOccupancyAfterMove(blackKing)
				move.setMoveType(castleKingside)
				moves = append(moves, move)
			}
		}
		if b.A8RookHasNeverMoved && (b.BlackPieces&b.Rooks&A8 != 0) {
			if (b.BlackPieces|b.WhitePieces)&(D8|C8|B8) == 0 && attackedSquares&(E8|D8|C8) == 0 {
				var move Move
				move.setOriginFromBB(E8) //all that's needed to update castling rights
				move.setOriginOccupancy(blackKing)
				move.setDestOccupancyBeforeMove(empty)
				move.setDestOccupancyAfterMove(blackKing)
				move.setMoveType(castleQueenside)
				moves = append(moves, move)
			}
		}
	}
	ch <- moves
}
