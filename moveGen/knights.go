package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

func (b *Board) getKnightDefendedSquares() (defendedSquares uint64) {
	var knights uint64
	var ownPieces uint64

	if b.IsWhiteMove {
		ownPieces = b.WhitePieces
		knights = b.Knights & b.WhitePieces
	} else {
		ownPieces = b.BlackPieces
		knights = b.Knights & b.BlackPieces
	}

	for bits.OnesCount64(knights) != 0 { //While there are still knights left
		currentSquare := utils.IsolateLsb(knights)
		defendedSquares |= KnightMask[bits.TrailingZeros64(currentSquare)] &^ ownPieces //get all squares the current knight can attack
		knights ^= currentSquare                                                        //clear the knight we just calculated
	}
	return defendedSquares
}

func (b *Board) getKnightMoves() []Move {
	var moves []Move
	var knights uint64
	var ownPieces uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalMove)

	if b.IsWhiteMove {
		ownPieces = b.WhitePieces
		knights = b.Knights & b.WhitePieces
		baseMove.setOriginOccupancy(whiteKnight)
		baseMove.setDestOccupancyAfterMove(whiteKnight)
	} else {
		ownPieces = b.BlackPieces
		knights = b.Knights & b.BlackPieces
		baseMove.setOriginOccupancy(blackKnight)
		baseMove.setDestOccupancyAfterMove(blackKnight)
	}

	for bits.OnesCount64(knights) != 0 { //While there are still knights left
		newMove := baseMove
		currentSquare := utils.IsolateLsb(knights)
		newMove.setOriginFromBB(currentSquare)
		possibleAttacks := KnightMask[bits.TrailingZeros64(currentSquare)] &^ ownPieces //get all squares the current knight can attack

		for possibleAttacks != 0 { //for every square that we can attack,
			move := newMove //copy our move with origin and move type
			attack := utils.IsolateLsb(possibleAttacks)
			move.setDestFromBB(attack)
			move.setDestOccupancyBeforeMove(b.getTileOccupancy(attack)) //note the piece (or lack of) that's on the square before we capture
			possibleAttacks ^= attack

			//Only add move if your king is not in check
			//Note that we don't pass if our own king is in check at start of move,
			//Because we'd have to check regardless. This is because you can put your own king in check if pinned.
			if b.IsWhiteMove { //todo: optimize
				undo := b.ApplyMove(move)
				//Must be attacked by self because ApplyMove flips the turn
				if b.GetSquaresAttackedThisHalfTurn()&(b.Kings&b.WhitePieces) != 0 { //If we are in check
					undo()
					continue
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
			//If we are not in check, add the move to legal moves
			moves = append(moves, move)
		}
		knights ^= currentSquare
	}
	return moves
}
