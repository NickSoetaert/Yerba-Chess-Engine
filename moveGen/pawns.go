package moveGen

import (
	"Yerba/utils"
	"fmt"
)

// TODO Don't pass by value - optimize
func (b Board) getPawnMoves(c chan []Move) {
	var moves []Move
	if b.IsWhiteMove {
		b.Pawns &= b.WhitePieces
	} else {
		b.Pawns &= b.BlackPieces
	}

	moves = append(moves, b.pawnNormalCaptures()...)
	moves = append(moves, b.pawnSinglePushMoves()...)
	moves = append(moves, b.pawnDoublePushMoves()...)
	moves = append(moves, b.enPassantCaptures()...)

	c <- moves
}

//Expects a pawn that is eligible for a promotion, and will return all possible promotions.
func pawnPromotionsHelper(move Move, isWhiteToMove bool) (allMoves []Move) {
	if isWhiteToMove {
		if (move.getDestSquare() & EighthRank) == 0 { //if there are no possible promotions, return base move
			panic(fmt.Sprintf("impossible promotion - dest square: %064b", move.getDestSquare()))
		}
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteKnight))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteBishop))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteRook))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteQueen))
	} else {
		if (move.getDestSquare() & FirstRank) == 0 {
			panic("impossible promotion")
		}
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(blackKnight))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(blackBishop))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(blackRook))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(blackQueen))
	}
	return allMoves
}

func (b Board) pawnSinglePushMoves() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalMove)
	baseMove.setDestOccupancyBeforeMove(empty)

	if b.IsWhiteMove {
		openSquares = (b.Pawns & b.WhitePieces) << 8
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares = (b.Pawns & b.BlackPieces) >> 8
		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	}

	openSquares &^= b.WhitePieces | b.BlackPieces //Filter out all squares with pieces on them
	for openSquares != 0 {                        //Convert all available squares to a Move
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)

		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 8) //record the square we started at

			if (dest & EighthRank) != 0 { // Factor in if a promotion is possible.
				for _, promotion := range pawnPromotionsHelper(newMove, true) {
					moves = append(moves, promotion)
				}
			} else { //If not, just add the unpromoted move to list of moves.
				moves = append(moves, newMove)
			}

		} else { //if black's move
			newMove.setOriginFromBB(dest << 8)

			if (dest & FirstRank) != 0 { // Factor in if a promotion is possible.
				for _, promotion := range pawnPromotionsHelper(newMove, false) {
					moves = append(moves, promotion)
				}
			} else { //If not, just add the unpromoted move to list of moves.
				moves = append(moves, newMove)
			}
		}
		openSquares &= openSquares - 1
	}
	return moves
}

func (b Board) pawnNormalCaptures() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalMove)

	if b.IsWhiteMove {
		openSquares = ((b.Pawns << 7) & ^HFile) & b.BlackPieces
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares = ((b.Pawns >> 7) & ^AFile) & b.WhitePieces
		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}

	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		newMove.setDestOccupancyBeforeMove(b.getTileOccupancy(dest))
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 7)
		} else {
			newMove.setOriginFromBB(dest << 7)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}

	if b.IsWhiteMove {
		openSquares |= ((b.Pawns << 9) & ^AFile) & b.BlackPieces
	} else {
		openSquares |= ((b.Pawns >> 9) & ^HFile) & b.WhitePieces
	}
	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)

		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 9)

			//Check for promotions
			if (dest & EighthRank) != 0 {
				for _, promotion := range pawnPromotionsHelper(newMove, true) {
					moves = append(moves, promotion)
				}
			} else { //If not, just add the unpromoted move to list of moves.
				moves = append(moves, newMove)
			}

		} else { //If black's move
			newMove.setOriginFromBB(dest << 9)

			if (dest & FirstRank) != 0 {
				for _, promotion := range pawnPromotionsHelper(newMove, false) {
					moves = append(moves, promotion)
				}
			} else {
				moves = append(moves, newMove)
			}
		}
		openSquares &= openSquares - 1
	}

	return moves
}

func (b Board) pawnDoublePushMoves() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(pawnDoublePush)
	baseMove.setDestOccupancyBeforeMove(empty)
	if b.IsWhiteMove {
		//Get moves that move forward 2 ranks and end up on proper rank, and don't jump over anything
		openSquares |= ((b.Pawns << 16) & FourthRank) ^ (((b.WhitePieces | b.BlackPieces) & ThirdRank) << 8)
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares |= ((b.Pawns >> 16) & FifthRank) ^ (((b.WhitePieces | b.BlackPieces) & SixthRank) >> 8)
		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}
	openSquares = openSquares & ^(b.WhitePieces | b.BlackPieces)

	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 16)
		} else {
			newMove.setOriginFromBB(dest << 16)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}
	return moves
}

func (b Board) enPassantCaptures() (moves []Move) {
	if b.EnPassantFile == 0 {
		return nil
	}
	baseMove := Move(0)
	baseMove.setMoveType(enPassantCapture)
	baseMove.setDestOccupancyBeforeMove(empty)
	var openSquares uint64

	if b.IsWhiteMove {
		openSquares |= ((b.Pawns << 7) & ^HFile) & (uint64(b.EnPassantFile) << 16)
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares |= ((b.Pawns >> 7) & ^AFile) & (uint64(b.EnPassantFile) << 40)
		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 7)
		} else {
			newMove.setOriginFromBB(dest << 7)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}

	if b.IsWhiteMove {
		openSquares |= ((b.Pawns << 9) & ^HFile) & (uint64(b.EnPassantFile) << 16)
	} else {
		openSquares |= ((b.Pawns >> 9) & ^AFile) & (uint64(b.EnPassantFile) << 40)
	}
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 9)
		} else {
			newMove.setOriginFromBB(dest << 9)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}
	return moves
}
