package moveGen

import (
	"Yerba/utils"
)

func GetPawnMoves(pawns, whitePieces, blackPieces uint64, isWhiteToMove bool, enPassantFile uint8) (moves []Move) {
	if isWhiteToMove {
		pawns &= whitePieces
	} else {
		pawns &= blackPieces
	}

	moves = append(moves, pawnNormalCaptures(pawns, whitePieces, blackPieces, isWhiteToMove)...)
	moves = append(moves, pawnSinglePushMoves(pawns, whitePieces, blackPieces, isWhiteToMove)...)
	moves = append(moves, pawnDoublePushMoves(pawns, whitePieces, blackPieces, isWhiteToMove)...)
	moves = append(moves, pawnEnPassantCaptures(pawns, whitePieces, blackPieces, isWhiteToMove, enPassantFile)...)

	return moves
}

func pawnSinglePushMoves(pawns, whitePieces, blackPieces uint64, isWhiteToMove bool) (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalPawnPush)
	if isWhiteToMove {
		openSquares = (pawns << 8) ^ EighthRank
	} else {
		openSquares = (pawns >> 8) ^ FirstRank
	}
	openSquares = openSquares & ^(whitePieces | blackPieces)

	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDest(dest)
		if isWhiteToMove{
			newMove.setOrigin(dest>>8)
		} else {
			newMove.setOrigin(dest<<8)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares-1
	}
	return moves
}

func pawnDoublePushMoves(pawns, whitePieces, blackPieces uint64, isWhiteToMove bool) (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(pawnDoublePush)
	if isWhiteToMove {
		//Get moves that move forward 2 ranks and end up on proper rank, and don't jump over anything
		openSquares |= ((pawns << 16) & FourthRank) ^ (((whitePieces|blackPieces) & ThirdRank) << 8)
	} else {
		openSquares |= ((pawns >> 16) & FifthRank ) ^ (((whitePieces|blackPieces) & SixthRank) >> 8)
	}
	openSquares = openSquares & ^(whitePieces | blackPieces)

	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDest(dest)
		if isWhiteToMove{
			newMove.setOrigin(dest>>16)
		} else {
			newMove.setOrigin(dest<<16)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares-1
	}
	return moves
}

func pawnNormalCaptures(pawns, whitePieces, blackPieces uint64, isWhiteToMove bool) (moves []Move) {
	var openSquares uint64
	baseMove := Move(0)
	baseMove.setMoveType(normalPawnCapture)
	if isWhiteToMove {
		openSquares = ((pawns << 7) & ^HFile) & blackPieces
	} else {
		openSquares = ((pawns >> 7) & ^AFile) & whitePieces
	}

	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDest(dest)
		if isWhiteToMove{
			newMove.setOrigin(dest>>7)
		} else {
			newMove.setOrigin(dest<<7)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares-1
	}

	if isWhiteToMove {
		openSquares |= ((pawns << 9) & ^AFile) & blackPieces
	} else {
		openSquares |= ((pawns >> 9) & ^HFile) & whitePieces
	}
	//Convert all available squares to a Move
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDest(dest)
		if isWhiteToMove{
			newMove.setOrigin(dest>>9)
		} else {
			newMove.setOrigin(dest<<9)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares-1
	}

	return moves
}

func pawnEnPassantCaptures(pawns, whitePieces, blackPieces uint64, isWhiteToMove bool, enPassantFile uint8) (moves []Move) {
	if enPassantFile == 0 {
		return nil
	}
	baseMove := Move(0)
	baseMove.setMoveType(enPassantCapture)
	var openSquares uint64

	if isWhiteToMove {
		openSquares |= ((pawns << 7) & ^HFile) & (uint64(enPassantFile) << 16)
	} else {
		openSquares |= ((pawns >> 7) & ^AFile) & (uint64(enPassantFile) << 40)
	}
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDest(dest)
		if isWhiteToMove{
			newMove.setOrigin(dest>>7)
		} else {
			newMove.setOrigin(dest<<7)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares-1
	}

	if isWhiteToMove {
		openSquares |= ((pawns << 9) & ^HFile) & (uint64(enPassantFile) << 16)
	} else {
		openSquares |= ((pawns >> 9) & ^AFile) & (uint64(enPassantFile) << 40)
	}
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDest(dest)
		if isWhiteToMove{
			newMove.setOrigin(dest>>9)
		} else {
			newMove.setOrigin(dest<<9)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares-1
	}

	return moves
}