package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

// Gets all squares that the king is attacking, regardless of if it can move there or not.
func (b *Board) getKingDefendedSquares() (defendedSquares uint64) {
	var currentSquare uint64
	if b.IsWhiteMove {
		currentSquare = utils.IsolateLsb(b.Kings & b.WhitePieces)
	} else {
		currentSquare = utils.IsolateLsb(b.Kings & b.BlackPieces)
	}

	return KingMask[bits.TrailingZeros64(currentSquare)]
}

// Gets all squares that a King can legally move to without castling this turn.
func (b Board) getNormalKingMoves(attackedSquares uint64, ch chan []Move) {
	var moves []Move
	//todo - remove this once done with debug
	//ch <- moves
	//return
	var currentSquare uint64
	var possibleAttacks uint64

	baseMove := Move(0)
	baseMove.setMoveType(normalMove)

	if b.IsWhiteMove {
		currentSquare = utils.IsolateLsb(b.Kings & b.WhitePieces)
		if bits.TrailingZeros64(currentSquare) >= 64 {
			utils.PrintBinaryBoard(b.Kings)
			PrintBoard(b)
		}
		possibleAttacks = KingMask[bits.TrailingZeros64(currentSquare)] &^ attackedSquares &^ b.WhitePieces
		baseMove.setOriginOccupancy(whiteKing)
		baseMove.setDestOccupancyAfterMove(whiteKing)
	} else {
		currentSquare = utils.IsolateLsb(b.Kings & b.BlackPieces)
		if bits.TrailingZeros64(currentSquare) >= 64 {
			utils.PrintBinaryBoard(b.Kings)
			PrintBoard(b)
		}
		possibleAttacks = KingMask[bits.TrailingZeros64(currentSquare)] &^ attackedSquares &^ b.BlackPieces
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

		//Only add attack if king isn't in check.
		//Say a rook is on A8, and king is on G8. This prevents king from thinking H8 is safe. Clearly room to optimize.
		if b.IsWhiteMove {
			undo := b.ApplyMove(move)
			//Must be attacked by self because ApplyMove flips the turn
			if b.GetSquaresAttackedThisHalfTurn()&(b.Kings&b.WhitePieces) != 0 { //If we are in check
				undo()
				continue //ignore this move because it is illegal
			}
			undo()
		} else {
			undo := b.ApplyMove(move)
			//Must be attacked by self because ApplyMove flips the turn
			if b.GetSquaresAttackedThisHalfTurn()&(b.Kings&b.BlackPieces) != 0 { //If we are in check
				undo()
				continue
			}
			undo()
		}

		moves = append(moves, move)
	}
	ch <- moves
}

func (b Board) getCastlingMoves(attackedSquares uint64, ch chan []Move) {
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
