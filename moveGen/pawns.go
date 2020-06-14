package moveGen

import (
	"Yerba/utils"
)

//TODO: promotions

// TODO
// TODO DANGER: I stopped half way through a refactor here. Setting b.Pawns &= b.WhitePieces will mess up the board state.
// TODO Either make pass by value (extra memory), do AND checks for each time you reference the pawns, or refactor to have
// TODO a larger Board object that has both White and Black pieces for each piece type.
func (b *Board) getPawnMoves(c chan []Move) {
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


func (b Board) allPossiblePromotionsHelper(move Move) (allPromotions []Move) {
	allPromotions = append(allPromotions, move.copyMoveSetType(knightPromotion))
	

	return nil
}

func (b Board) pawnPromotionsHelper() (moves []Move) {

	if b.IsWhiteMove {
		b.Pawns &= SeventhRank //If it is white to move, the only pawns that can promote are on the 7th rank.
	} else {
		b.Pawns &= SecondRank
	}

	return nil
}


func (b Board) pawnSinglePushMoves() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalPawnPush)
	if b.IsWhiteMove {
		//ensure we don't push b.Pawns to the 8th rank
		openSquares = (b.Pawns << 8) ^ EighthRank
	} else {
		openSquares = (b.Pawns >> 8) ^ FirstRank
	}
	openSquares = openSquares & ^(b.WhitePieces | b.BlackPieces)

	//Convert all available squares to a Move
	for openSquares != 0 {
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
	return moves
}

func (b Board) pawnNormalCaptures() (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalPawnCapture)
	if b.IsWhiteMove {
		openSquares = ((b.Pawns << 7) & ^HFile) & b.BlackPieces
	} else {
		openSquares = ((b.Pawns >> 7) & ^AFile) & b.WhitePieces
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

	return moves
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
