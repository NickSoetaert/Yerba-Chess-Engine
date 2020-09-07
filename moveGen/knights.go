package moveGen

import (
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
		originSquareMove := baseMove
		currentSquare := uint8(bits.TrailingZeros64(knights))     //square number that we're looking at
		originSquareMove.setOriginFromSquare(currentSquare)       //set this move to start at said square
		possibleAttacks := KnightMask[currentSquare] &^ ownPieces //get all squares the current knight can attack

		for possibleAttacks != 0 { //for every square that we can attack,
			move := originSquareMove //copy our move with origin and move type
			attack := uint8(bits.TrailingZeros64(possibleAttacks))
			move.setDestFromSquare(attack)
			moves = append(moves, move)
			possibleAttacks ^= uint64(1 << attack)
		}
		knights ^= uint64(1 << currentSquare) //now clear that knight for the next loop iteration
	}
	ch <- moves
}
