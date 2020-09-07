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
	fmt.Printf("pawn captures: %v\n", len(moves))
	moves = append(moves, b.pawnSinglePushMoves()...)
	fmt.Printf("pawn single push: %v\n", len(moves))
	moves = append(moves, b.pawnDoublePushMoves()...)
	fmt.Printf("pawn double push: %v\n", len(moves))
	moves = append(moves, b.enPassantCaptures()...)
	fmt.Printf("pawn ep: %v\n", len(moves))

	c <- moves
}


//Expects a pawn that is eligible for a promotion, and will return all possible promotions.
func pawnPromotionsHelper(move Move, isWhiteToMove bool) (allMoves []Move) {
	if isWhiteToMove {
		if (move.getDestSquare() & EighthRank) == 0 { //if there are no possible promotions, return base move
			return nil
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
	openSquares = b.Pawns << 8
	openSquares ^= b.WhitePieces | b.BlackPieces //Get all squares without a piece on them

	for openSquares != 0 { 	//Convert all available squares to a Move

		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 8)
		} else {
			newMove.setOriginFromBB(dest << 8)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}
	fmt.Printf("pawn pre-promotion: %v", len(moves))
	//Todo - optimize by only applying to moves on last rank
	var promotedMoves []Move
	for _, move := range moves {
		if move.
		for _, promotion := range pawnPromotionsHelper(move, b.IsWhiteMove) {
			promotedMoves = append(moves, promotion)
		}
	}
	fmt.Printf("pawn post-promotion: %v", len(promotedMoves))
	return promotedMoves
}

func (b Board) pawnNormalCaptures() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalMove)

	if b.IsWhiteMove {
		//openSquares = ((b.Pawns << 7) & ^HFile) & b.BlackPieces //old
		openSquares = ((b.Pawns & b.WhitePieces) << 7) & b.BlackPieces //all squares that white pawns can attack
	} else {
		//openSquares = ((b.Pawns >> 7) & ^AFile) & b.WhitePieces //old
		openSquares = ((b.Pawns & b.BlackPieces) >> 7) & b.WhitePieces //all squares that black pawns can attack
	}

	//Convert all available squares to a Move
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
		} else {
			newMove.setOriginFromBB(dest << 9)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}

	//Todo - optimize by only applying to moves on last rank
	var promotedMoves []Move
	for _, move := range moves {
		for _, promotion := range pawnPromotionsHelper(move, b.IsWhiteMove) {
			promotedMoves = append(moves, promotion)
		}
	}
	return promotedMoves
}

func (b Board) pawnDoublePushMoves() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(pawnDoublePush)
	if b.IsWhiteMove {
		//Get moves that move forward 2 ranks and end up on proper rank, and don't jump over anything
		openSquares |= ((b.Pawns << 16) & FourthRank) ^ (((b.WhitePieces | b.BlackPieces) & ThirdRank) << 8)
	} else {
		openSquares |= ((b.Pawns >> 16) & FifthRank) ^ (((b.WhitePieces | b.BlackPieces) & SixthRank) >> 8)
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
	var openSquares uint64

	if b.IsWhiteMove {
		openSquares |= ((b.Pawns << 7) & ^HFile) & (uint64(b.EnPassantFile) << 16)
	} else {
		openSquares |= ((b.Pawns >> 7) & ^AFile) & (uint64(b.EnPassantFile) << 40)
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
