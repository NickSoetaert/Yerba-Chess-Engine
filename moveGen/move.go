package moveGen

import (
	"Yerba/utils"
	"fmt"
	"math/bits"
)

type UndoMove func()

//Given a starting board and a move, ApplyMove applies said move to said board,
//and returns a function that will undo the previously applied move.
//ApplyMove does NOT check for legality; that is the responsibility of MoveGen.
func (b *Board) ApplyMove(m Move) UndoMove {
	oldBoard := *b //TODO: Optimize

	b.clearTargetSquare(m) //Must clear starting square for all move types

	switch m.getMoveType() {
	case normalMove:
		b.putPieceOnTargetSquare(m)

	case pawnDoublePush:
		b.EnPassantFile = uint8(bits.TrailingZeros64(m.getOriginSquare()) % 8) //can capture e.p. next turn
		b.putPieceOnTargetSquare(m)

	case enPassantCapture:
		//add and remove pawns for Pawns bitboards
		b.Pawns |= m.getDestSquare()                                               //add pawn
		b.Pawns = b.Pawns &^ enPassantFileToSquare(b.EnPassantFile, b.IsWhiteMove) //remove pawn

		if b.IsWhiteMove { //add and remove pawns for WhitePieces and BlackPieces bitboards
			b.WhitePieces |= m.getDestSquare()                                                     //add pawn
			b.BlackPieces = b.BlackPieces &^ enPassantFileToSquare(b.EnPassantFile, b.IsWhiteMove) //remove pawn
		} else {
			b.BlackPieces |= m.getDestSquare()                                                     //add pawn
			b.WhitePieces = b.WhitePieces &^ enPassantFileToSquare(b.EnPassantFile, b.IsWhiteMove) //remove pawn
		}

	case castleKingside:
		if b.IsWhiteMove {
			b.Kings |= G1
			b.Rooks = b.Rooks &^ H1
			b.WhitePieces = b.WhitePieces &^ A1
			b.Rooks |= F1
		} else {
			b.Kings |= G8
			b.Rooks = b.Rooks &^ H8
			b.BlackPieces = b.BlackPieces &^ A8
			b.Rooks |= F8
		}

	case castleQueenside:
		if b.IsWhiteMove {
			b.Kings |= C1
			b.Rooks = b.Rooks &^ A1
			b.WhitePieces = b.WhitePieces &^ A1
			b.Rooks |= D1
			b.WhitePieces |= D1
		} else {
			b.Kings |= C8
			b.Rooks = b.Rooks &^ A8
			b.BlackPieces = b.BlackPieces &^ A8
			b.Rooks |= D8
			b.BlackPieces |= D8
		}
	default:
		panic("didn't set move type for move")
	}

	b.updateCastlingRights(m) //update flags relevant to castling

	//Clear en passant capture unless we just pushed a pawn two squares
	if m.getMoveType() != pawnDoublePush {
		b.EnPassantFile = 0
	}

	//Finally, change who's turn it is
	b.IsWhiteMove = !b.IsWhiteMove

	return func() { //Idea here is to return a function that will undo this move.
		//todo
		*b = oldBoard
	}
}

func (b *Board) updateCastlingRights(m Move) {
	//We don't need to check for piece type because once that piece is off its home square,
	//its corresponding bool is already false. Hence no harm in setting a false value to false.
	if m.getOriginSquare() == A1 {
		b.A1RookHasNeverMoved = false
	} else if m.getOriginSquare() == A8 {
		b.A8RookHasNeverMoved = false
	} else if m.getOriginSquare() == H1 {
		b.H1RookHasNeverMoved = false
	} else if m.getOriginSquare() == H8 {
		b.H8RookHasNeverMoved = false
	} else if m.getOriginSquare() == E1 {
		b.WhiteKingHasNeverMoved = false
	} else if m.getOriginSquare() == E8 {
		b.BlackKingHasNeverMoved = false
	}
}

//Given a file number (with 1==A, 8==h) and who's turn it is,
//Returns the square on which an en passant capture would take place. Wordy but fast.
func enPassantFileToSquare(file uint8, isWhiteToMove bool) uint64 {
	if isWhiteToMove {
		switch file {
		case 1:
			return A4
		case 2:
			return B4
		case 3:
			return C4
		case 4:
			return D4
		case 5:
			return E4
		case 6:
			return F4
		case 7:
			return G4
		case 8:
			return H4
		default:
			panic(fmt.Sprintf("impossible e.p. file %v", file))
		}
	} else {
		switch file {
		case 1:
			return A5
		case 2:
			return B5
		case 3:
			return C5
		case 4:
			return D5
		case 5:
			return E5
		case 6:
			return F5
		case 7:
			return G5
		case 8:
			return H5
		default:
			panic(fmt.Sprintf("impossible e.p. file %v", file))
		}
	}
}

// Removes a piece from its start square
func (b *Board) clearOriginSquare(m Move) {
	if b.IsWhiteMove {
		b.WhitePieces = b.WhitePieces &^ m.getOriginSquare()
	} else {
		b.BlackPieces = b.BlackPieces &^ m.getOriginSquare()
	}

	switch m.getOriginOccupancy() {
	case whitePawn:
		b.Pawns = b.Pawns &^ m.getOriginSquare()
	case whiteKnight:
		b.Knights = b.Knights &^ m.getOriginSquare()
	case whiteBishop:
		b.Bishops = b.Bishops &^ m.getOriginSquare()
	case whiteRook:
		b.Rooks = b.Rooks &^ m.getOriginSquare()
	case whiteQueen:
		b.Queens = b.Queens &^ m.getOriginSquare()
	case whiteKing:
		b.Kings = b.Kings &^ m.getOriginSquare()

	case blackPawn:
		b.Pawns = b.Pawns &^ m.getOriginSquare()
	case blackKnight:
		b.Knights = b.Knights &^ m.getOriginSquare()
	case blackBishop:
		b.Bishops = b.Bishops &^ m.getOriginSquare()
	case blackRook:
		b.Rooks = b.Rooks &^ m.getOriginSquare()
	case blackQueen:
		b.Queens = b.Queens &^ m.getOriginSquare()
	case blackKing:
		b.Kings = b.Kings &^ m.getOriginSquare()
	default:
		panic(fmt.Sprintf("m.getOriginOccupancy() returned: %v", m.getOriginOccupancy()))
	}
}

//Captures (removes) the piece on the target square
func (b *Board) clearTargetSquare(m Move) {
	if b.IsWhiteMove {
		b.WhitePieces = b.WhitePieces &^ m.getDestSquare()
	} else {
		b.BlackPieces = b.BlackPieces &^ m.getDestSquare()
	}

	switch m.getDestOccupancyBeforeMove() {
	case whitePawn:
		b.Pawns = b.Pawns &^ m.getDestSquare()
	case whiteKnight:
		b.Knights = b.Knights &^ m.getDestSquare()
	case whiteBishop:
		b.Bishops = b.Bishops &^ m.getDestSquare()
	case whiteRook:
		b.Rooks = b.Rooks &^ m.getDestSquare()
	case whiteQueen:
		b.Queens = b.Queens &^ m.getDestSquare()
	case whiteKing:
		b.Kings = b.Kings &^ m.getDestSquare()

	case blackPawn:
		b.Pawns = b.Pawns &^ m.getDestSquare()
	case blackKnight:
		b.Knights = b.Knights &^ m.getDestSquare()
	case blackBishop:
		b.Bishops = b.Bishops &^ m.getDestSquare()
	case blackRook:
		b.Rooks = b.Rooks &^ m.getDestSquare()
	case blackQueen:
		b.Queens = b.Queens &^ m.getDestSquare()
	case blackKing:
		b.Kings = b.Kings &^ m.getDestSquare()
	case empty:
		//do nothing
	default:
		panic(fmt.Sprintf("m.getDestOccupancyBeforeMove() returned %v - piece %064b", m.getDestOccupancyBeforeMove(), m.getOriginSquare()))
	}
}

//Puts the piece on its destination square, accounting for promotions.
func (b *Board) putPieceOnTargetSquare(m Move) {
	if b.IsWhiteMove {
		b.WhitePieces |= m.getDestSquare()
	} else {
		b.BlackPieces |= m.getDestSquare()
	}

	switch m.getDestOccupancyAfterMove() {
	case whitePawn:
		b.Pawns |= m.getDestSquare()
	case whiteKnight:
		b.Knights |= m.getDestSquare()
	case whiteBishop:
		b.Bishops |= m.getDestSquare()
	case whiteRook:
		b.Rooks |= m.getDestSquare()
	case whiteQueen:
		b.Queens |= m.getDestSquare()
	case whiteKing:
		b.Kings |= m.getDestSquare()

	case blackPawn:
		b.Pawns |= m.getDestSquare()
	case blackKnight:
		b.Knights |= m.getDestSquare()
	case blackBishop:
		b.Bishops |= m.getDestSquare()
	case blackRook:
		b.Rooks |= m.getDestSquare()
	case blackQueen:
		b.Queens |= m.getDestSquare()
	case blackKing:
		b.Kings |= m.getDestSquare()
	default:
		panic(fmt.Sprintf("m.getDestOccupancyAfterMove() returned: %v, piece:%032b\n", m.getDestOccupancyAfterMove(), m))
	}
}

//Returns binary-board (64 bit) representation of origin square.
//Origin square: bits 0-5
func (m Move) getOriginSquare() uint64 {
	return 1 << (m >> 26)
}

//Returns binary-board (64 bit) representation of destination square.
//Destination square: bits 6-11
func (m Move) getDestSquare() uint64 {
	fmt.Printf("%032b\n",m)
	fmt.Printf("unshifted dest:        %b\n",utils.IsolateBitsU32(uint32(m), destSquareBitsStart, destSquareBitsEnd))
	return 1 << uint64(utils.IsolateBitsU32(uint32(m), destSquareBitsStart, destSquareBitsEnd))
}

//Returns a tileOccupancy representing the piece that occupied the origin square
//Orig occupancy bits: 12-15
func (m Move) getOriginOccupancy() tileOccupancy {
	return tileOccupancy(utils.IsolateBitsU32(uint32(m), originSquareOccBitsStart, originSquareOccBitsEnd))
}

//returns the tileOccupancy that was on the destination square before the move was made.
//bits: 16-20
func (m Move) getDestOccupancyBeforeMove() tileOccupancy {
	return tileOccupancy(utils.IsolateBitsU32(uint32(m), destSquarePreMoveOccBitsStart, destSquarePreMoveOccBitsEnd))
}

//returns the tileOccupancy that is on the destination square after the move was made.
//bits: 21-25
func (m Move) getDestOccupancyAfterMove() tileOccupancy {
	return tileOccupancy(utils.IsolateBitsU32(uint32(m), destSquarePostMoveOccBitsStart, destSquarePostMoveOccBitsEnd))
}

//Returns the moveType bits (bits 26-29)
func (m Move) getMoveType() moveType {
	return moveType(utils.IsolateBitsU32(uint32(m), moveTypeBitsStart, moveTypeBitsEnd))
}

//Expects a bitboard with a pop count of one, and sets the origin square of given move to that square.
//Origin square coordinate bits: 0-5
func (m *Move) setOriginFromBB(origin uint64) {
	*m = *m &^ (0b111111 << moveTypeBitsStart)                    //clear origin bits
	*m |= Move(bits.TrailingZeros64(origin)) << moveTypeBitsStart //set the cleared bits
}

//Expects a square position - e.g. 7 to represent A8, and will set the origin square to be that square.
func (m *Move) setOriginFromSquare(origin uint8) {
	*m = *m &^ (0b111111 << moveTypeBitsStart) //clear origin bits
	*m |= Move(origin) << moveTypeBitsStart    //set the cleared bits
}

//Expects a bitboard with a pop count of one, and sets the destination square of given move to that square.
func (m *Move) setDestFromBB(dest uint64) {
	//fmt.Printf("move before:%032b\n",*m)
	//fmt.Printf("trailing 0s: %v\n",bits.TrailingZeros64(dest))
	*m = Move(utils.SetBitsU32(uint32(*m), destSquareBitsStart, destSquareBitsEnd, uint32(bits.TrailingZeros64(dest))))
	//fmt.Printf("move after: %032b\n",*m)
}

func (m *Move) setDestFromSquare(dest uint8) {
	*m = Move(utils.SetBitsU32(uint32(*m), destSquareBitsStart, destSquareBitsEnd, uint32(dest)))
}

//Move type bits: 26-29
func (m *Move) setMoveType(mt moveType) {
	*m = Move(utils.SetBitsU32(uint32(*m), moveTypeBitsStart, moveTypeBitsEnd, uint32(mt)))
}

//Duplicates a Move, but also sets the moveType
func (m *Move) copyMoveAndSetType(mt moveType) Move {
	newMove := *m
	newMove.setMoveType(mt)
	return newMove
}

//Sets the occupancy (piece type) of the origin square.
//You can think of this as the piece you're picking up.
func (m *Move) setOriginOccupancy(oldPiece tileOccupancy) {
	*m = Move(utils.SetBitsU32(uint32(*m), originSquareOccBitsStart, originSquareOccBitsEnd, uint32(oldPiece)))
}

//Sets the occupancy (piece type) of the destination square.
func (m *Move) setDestOccupancyAfterMove(newPiece tileOccupancy) {
	*m = Move(utils.SetBitsU32(uint32(*m), destSquarePostMoveOccBitsStart, destSquarePostMoveOccBitsEnd, uint32(newPiece)))
}

//Sets the piece that was on the target square before the move
func (m *Move) setDestOccupancyBeforeMove(oldPiece tileOccupancy) {
	*m = Move(utils.SetBitsU32(uint32(*m), destSquarePreMoveOccBitsStart, destSquarePreMoveOccBitsEnd, uint32(oldPiece)))
}

//Helper for pawn promotions. Copies a move, but also sets the occupancy of the destination square.
func (m *Move) copyMoveAndSetDestOccupancy(newPiece tileOccupancy) Move {
	newMove := *m
	newMove.setDestOccupancyAfterMove(newPiece)
	return newMove
}


