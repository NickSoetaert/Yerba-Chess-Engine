package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

func (b *Board) getKnightMoves(ch chan []Move) {
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

			moves = append(moves, move)
			possibleAttacks ^= attack
		}
		knights ^= currentSquare
	}
	ch <- moves
}
