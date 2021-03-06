package moveGen

import (
	"Yerba/utils"
)

func (b Board) GetPawnDefendedSquares() (defendedSquares uint64) {
	if b.IsWhiteMove {
		defendedSquares |= (b.Pawns & b.WhitePieces << 7) & ^HFile
		defendedSquares |= (b.Pawns & b.WhitePieces << 9) & ^AFile

	} else {
		defendedSquares |= (b.Pawns & b.BlackPieces >> 7) & ^AFile
		defendedSquares |= (b.Pawns & b.BlackPieces >> 9) & ^HFile
	}
	return defendedSquares
}

func (b Board) getPawnMoves() []Move {
	var unfilteredMoves []Move

	//If there are no pawns, take a shortcut and return no moves.
	if b.IsWhiteMove && (b.Pawns&b.WhitePieces == 0) {
		return unfilteredMoves
	}
	if !b.IsWhiteMove && (b.Pawns&b.BlackPieces == 0) {
		return unfilteredMoves
	}

	unfilteredMoves = append(unfilteredMoves, b.pawnNormalCaptures()...)
	unfilteredMoves = append(unfilteredMoves, b.pawnSinglePushMoves()...)
	unfilteredMoves = append(unfilteredMoves, b.pawnDoublePushMoves()...)
	unfilteredMoves = append(unfilteredMoves, b.enPassantCaptures()...)

	//Moves filtered by king being in check or not
	var filteredMoves []Move

	//Check if each move is legal because of king being in check
	if b.IsWhiteMove {
		for _, move := range unfilteredMoves {
			undo := b.ApplyMove(move)
			//Must be attacked by self because ApplyMove flips the turn
			if b.GetSquaresAttackedThisHalfTurn()&(b.Kings&b.WhitePieces) != 0 { //If we are in check
				undo()
				continue
			}
			undo()
			filteredMoves = append(filteredMoves, move)
		}
	} else {
		for _, move := range unfilteredMoves {
			undo := b.ApplyMove(move)
			//Must be attacked by self because ApplyMove flips the turn
			if b.GetSquaresAttackedThisHalfTurn()&(b.Kings&b.BlackPieces) != 0 { //If we are in check
				undo()
				continue
			}
			undo()
			filteredMoves = append(filteredMoves, move)
		}
	}

	return filteredMoves
}

//Expects a pawn move, and will return all possible promotions. If no promotions are possible, returns given move.
func pawnPromotionsHelper(move Move, isWhiteToMove bool) (allMoves []Move) {
	if isWhiteToMove {
		if (move.getDestSquare() & EighthRank) == 0 { //if there are no possible promotions, return default move
			return []Move{move}
		}
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteKnight))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteBishop))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteRook))
		allMoves = append(allMoves, move.copyMoveAndSetDestOccupancy(whiteQueen))
	} else {
		if (move.getDestSquare() & FirstRank) == 0 {
			return []Move{move}
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
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}

	openSquares &^= b.WhitePieces | b.BlackPieces //Filter out all squares with pieces on them
	for openSquares != 0 {                        //Convert all available squares to a Move
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)

		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 8) //record the square we started at
			for _, promotion := range pawnPromotionsHelper(newMove, true) {
				moves = append(moves, promotion)
			}
		} else { //if black's move
			newMove.setOriginFromBB(dest << 8)
			for _, promotion := range pawnPromotionsHelper(newMove, false) {
				moves = append(moves, promotion)
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

	// First, get all captures that go towards the H file
	if b.IsWhiteMove {
		openSquares |= ((b.Pawns & b.WhitePieces << 9) & ^AFile) & b.BlackPieces
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares = ((b.Pawns & b.BlackPieces >> 7) & ^AFile) & b.WhitePieces
		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}

	//Convert all available squares to a Move, and apply promotions if applicable.
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		newMove.setDestOccupancyBeforeMove(b.getTileOccupancy(dest))
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 9)
			for _, promotion := range pawnPromotionsHelper(newMove, true) {
				moves = append(moves, promotion)
			}

		} else { //If black's move
			newMove.setOriginFromBB(dest << 7)
			for _, promotion := range pawnPromotionsHelper(newMove, false) {
				moves = append(moves, promotion)
			}
		}
		openSquares &= openSquares - 1
	}

	//Second, get all captures towards the A file
	if b.IsWhiteMove {
		openSquares = ((b.Pawns & b.WhitePieces << 7) & ^HFile) & b.BlackPieces
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares |= ((b.Pawns & b.BlackPieces >> 9) & ^HFile) & b.WhitePieces
		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}

	//Convert all available squares to a Move, and apply promotions if applicable.
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		newMove.setDestOccupancyBeforeMove(b.getTileOccupancy(dest))
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 7)
			for _, promotion := range pawnPromotionsHelper(newMove, true) {
				moves = append(moves, promotion)
			}

		} else { //If black's move
			newMove.setOriginFromBB(dest << 9)

			for _, promotion := range pawnPromotionsHelper(newMove, false) {
				moves = append(moves, promotion)
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
		openSquares |= ((b.Pawns & b.WhitePieces << 16) & FourthRank) &^ (((b.WhitePieces | b.BlackPieces) & ThirdRank) << 8)
		//utils.PrintBinaryBoard(openSquares)
		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		openSquares |= ((b.Pawns & b.BlackPieces >> 16) & FifthRank) &^ (((b.WhitePieces | b.BlackPieces) & SixthRank) >> 8)
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

	//First, get all attacks that go towards the A file
	if b.IsWhiteMove {
		//get squares to left diagonal that can be attacked - avoid wrapping to h file
		//Do not get right diagonals yet, as we won't know the origin square
		reachable := ((b.Pawns & b.WhitePieces) << 7) &^ HFile

		//get squares that can actually be attacked, given current E.P. square.
		canEpCapture := enPassantFileToAttackedSquare(b.EnPassantFile, b.IsWhiteMove)

		openSquares |= reachable & canEpCapture //should only be 1 square

		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {

		//get squares to left diagonal that can be attacked - avoid wrapping to h file
		//Do not get right diagonals yet, as we won't know the origin square
		reachable := ((b.Pawns & b.BlackPieces) >> 9) &^ HFile

		//get squares that can actually be attacked, given current E.P. square.
		canEpCapture := enPassantFileToAttackedSquare(b.EnPassantFile, b.IsWhiteMove)

		openSquares |= reachable & canEpCapture //should only be 1 square

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
			newMove.setOriginFromBB(dest << 9)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}

	//Second, get all attacks that go towards the H file
	if b.IsWhiteMove {
		//get squares to left diagonal that can be attacked - avoid wrapping to A file
		reachable := ((b.Pawns & b.WhitePieces) << 9) &^ AFile

		//get squares that can actually be attacked, given current E.P. square.
		canEpCapture := enPassantFileToAttackedSquare(b.EnPassantFile, b.IsWhiteMove)

		openSquares |= reachable & canEpCapture //should only be 1 square

		baseMove.setOriginOccupancy(whitePawn)
		baseMove.setDestOccupancyAfterMove(whitePawn)
	} else {
		//get squares to left diagonal that can be attacked - avoid wrapping to A file
		reachable := ((b.Pawns & b.BlackPieces) >> 7) &^ AFile

		//get squares that can actually be attacked, given current E.P. square.
		canEpCapture := enPassantFileToAttackedSquare(b.EnPassantFile, b.IsWhiteMove)

		openSquares |= reachable & canEpCapture //should only be 1 square

		baseMove.setOriginOccupancy(blackPawn)
		baseMove.setDestOccupancyAfterMove(blackPawn)
	}
	//utils.PrintBinaryBoard(openSquares)
	for openSquares != 0 {
		dest := utils.IsolateLsb(openSquares)
		newMove := baseMove
		newMove.setDestFromBB(dest)
		if b.IsWhiteMove {
			newMove.setOriginFromBB(dest >> 9)
		} else {
			newMove.setOriginFromBB(dest << 7)
		}
		moves = append(moves, newMove)
		openSquares &= openSquares - 1
	}
	return moves
}
