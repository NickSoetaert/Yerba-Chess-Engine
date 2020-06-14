package moveGen

import (
	"Yerba/utils"
	"fmt"
	"math/bits"
)

//type Move: 32 bit unsigned int.
//Origin square coordinate bits: 0-5 (0 represents A1, and 64 represents H8)
//Destination square coordinate bits: 6-11
//Origin square occupancy type: 12-15
//Destination square occupancy type PRE MOVE: 16-20
//Destination square occupancy POST MOVE: 21-25
//Move type bits (needed for undo move function): 26-29
//Who's move: bit 30. (0 means White made the move, 1 means Black made the move.)
//bit 31: currently unused


//Move type bits
//0000 - normal move
//0001 - double pawn push (Needed for en passant)
//0010 - en passant capture
//0011 - move that takes away White's kingside castle rights (note you can lose castle rights from moving or a rook being captured)
//0100 - move that takes away White's queenside castle rights
//0101 - move that takes away Black's kingside castle rights
//0110 - move that takes away Black's kingside castle rights
//0111 - castle kingside
//1000 - castle queenside

type Move uint32

type UndoMove func()

//Given a starting board and a move, return the resulting board.
//Returns a function that undoes the previously applied move.
//ApplyMove does NOT check for legality; that is the responsibility of MoveGen.
func (b *Board) ApplyMove(m Move) UndoMove {
	oldBoard := *b //TODO: Optimize

	//Note - adding/removing pieces from the White/Black bitboards is taken care of in a single case after the switch.
	switch m.getMoveType() {
	case normalPawnPush:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.Pawns |= m.getDestSquare()             //add pawn

	case normalPawnCapture:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.clearTargetSquare(m.getDestSquare())   //clear captured piece
		b.Pawns |= m.getDestSquare()             //add pawn

	case pawnDoublePush:
		b.EnPassantFile = uint8(bits.TrailingZeros64(m.getOriginSquare()) % 8)
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.Pawns |= m.getDestSquare()             //add pawn

	case enPassantCapture:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.Pawns |= m.getDestSquare()             //add pawn
		if b.IsWhiteMove {
			b.clearTargetSquare(uint64(b.EnPassantFile) << 16) //remove captured piece
		} else {
			b.clearTargetSquare(uint64(b.EnPassantFile) << 40) //TODO: check
		}

	case knightMove:
		b.Knights = b.Knights &^ m.getOriginSquare()
		b.clearTargetSquare(m.getDestSquare())
		b.Knights |= m.getDestSquare()

	case bishopMove:
		b.Bishops = b.Bishops &^ m.getOriginSquare()
		b.clearTargetSquare(m.getDestSquare())
		b.Bishops |= m.getDestSquare()

	case rookMove:
		//It's possible that we were moving a non-original rook from a corner square, but that doesn't matter as
		//the original rook is no longer on that corner square. (So castle rights were revoked in the first place.)
		if m.getOriginSquare() == A1 {
			b.WhiteQueensideCastleRights = false
		} else if m.getOriginSquare() == H1 {
			b.WhiteKingsideCastleRights = false
		} else if m.getOriginSquare() == A8 {
			b.BlackQueensideCastleRights = false
		} else if m.getOriginSquare() == H8 {
			b.BlackKingsideCastleRights = false
		}

		b.Rooks = b.Rooks &^ m.getOriginSquare()
		b.clearTargetSquare(m.getDestSquare())
		b.Rooks |= m.getDestSquare()

	case queenMove:
		b.Queens = b.Queens &^ m.getOriginSquare()
		b.clearTargetSquare(m.getDestSquare())
		b.Queens |= m.getDestSquare()

	case normalKingMove:
		if b.IsWhiteMove {
			b.WhiteKingsideCastleRights = false
			b.WhiteQueensideCastleRights = false
		} else {
			b.BlackKingsideCastleRights = false
			b.WhiteQueensideCastleRights = false
		}
		b.Kings = b.Kings &^ m.getOriginSquare()
		b.clearTargetSquare(m.getDestSquare())
		b.Kings |= m.getDestSquare()

	case castleKingside:
		b.Kings = b.Kings &^ m.getOriginSquare()
		if b.IsWhiteMove {
			b.Kings |= G1
			b.Rooks = b.Rooks &^ H1
			b.Rooks |= F1
			b.WhiteKingsideCastleRights = false
			b.WhiteQueensideCastleRights = false
		} else {
			b.Kings |= G8
			b.Rooks = b.Rooks &^ H8
			b.Rooks |= F8
			b.BlackKingsideCastleRights = false
			b.BlackQueensideCastleRights = false
		}

	case castleQueenside:
		b.Kings = b.Kings &^ m.getOriginSquare()
		if b.IsWhiteMove {
			b.Kings |= C1
			b.Rooks = b.Rooks &^ A1
			b.WhitePieces = b.WhitePieces &^ A1
			b.Rooks |= D1
			b.WhitePieces |= D1
			b.WhiteKingsideCastleRights = false
			b.WhiteQueensideCastleRights = false
		} else {
			b.Kings |= C8
			b.Rooks = b.Rooks &^ A8
			b.BlackPieces = b.BlackPieces &^ A8
			b.Rooks |= D8
			b.BlackPieces |= D8
			b.BlackKingsideCastleRights = false
			b.BlackQueensideCastleRights = false
		}

	case knightPromotion:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.clearTargetSquare(m.getDestSquare())   //clear possible capture
		b.Knights |= m.getDestSquare()           //add knight

	case bishopPromotion:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.clearTargetSquare(m.getDestSquare())   //clear possible capture
		b.Bishops |= m.getDestSquare()           //add bishop

	case rookPromotion:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.clearTargetSquare(m.getDestSquare())   //clear possible capture
		b.Rooks |= m.getDestSquare()             //add rook

	case queenPromotion:
		b.Pawns = b.Pawns &^ m.getOriginSquare() //clear pawn
		b.clearTargetSquare(m.getDestSquare())   //clear possible capture
		b.Queens |= m.getDestSquare()            //add queen

	}

	//Clear en passant capture unless we just pushed a pawn two squares
	if m.getMoveType() != pawnDoublePush {
		b.EnPassantFile = 0
	}

	//Take care of adding/removing pieces from the color bitboards.
	if b.IsWhiteMove {
		b.WhitePieces = b.WhitePieces &^ m.getOriginSquare()
		b.WhitePieces |= m.getDestSquare()
	} else {
		b.BlackPieces = b.BlackPieces &^ m.getOriginSquare()
		b.BlackPieces |= m.getDestSquare()
	}

	//Finally, change who's turn it is
	b.IsWhiteMove = !b.IsWhiteMove

	return func() {
		*b = oldBoard
	}
}

//Removes all possible pieces in the way of a potential capture square
func (b *Board) clearTargetSquare(square uint64) {
	b.Pawns = b.Pawns &^ square
	b.Knights = b.Knights &^ square
	b.Bishops = b.Bishops &^ square
	b.Rooks = b.Rooks &^ square
	b.Queens = b.Queens &^ square

	if b.IsWhiteMove {
		b.BlackPieces = b.BlackPieces &^ square
	} else {
		b.WhitePieces = b.WhitePieces &^ square
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
	return 1 << uint64(utils.IsolateBitsU32(uint32(m), 6, 11))
}

//Returns a tileOccupancy representing the piece that occupied the origin square
//Orig occupancy bits: 12-15
func (m Move) getOriginOccupancy() tileOccupancy {
	return tileOccupancy(utils.IsolateBitsU32(uint32(m), 12, 15))
}

//returns the tileOccupancy that was on the destination square before the move was made.
//bits: 16-20
func (m Move) getDestOccupancyBeforeMove() tileOccupancy {
	return tileOccupancy(utils.IsolateBitsU32(uint32(m), 16, 20))
}

//returns the tileOccupancy that is on the destination square after the move was made.
//bits: 21-25
func (m Move) getDestOccupancyAfterMove() tileOccupancy {
	return tileOccupancy(utils.IsolateBitsU32(uint32(m), 21, 25))
}

//Returns the moveType bits (bits 26-29)
func (m Move) getMoveType() moveType {
	return moveType(utils.IsolateBitsU32(uint32(m), 26, 29))
}

//Expects a bitboard with a pop count of one, and sets the origin square of given move to that square.
func (m *Move) setOriginFromBB(origin uint64) {
	*m = *m &^ (0b111111 << 26)                    //clear origin bits
	*m |= Move(bits.TrailingZeros64(origin)) << 26 //set the cleared bits
}

//Expects a square position - e.g. 7 to represent A8, and will set the origin square to be that square.
func (m *Move) setOriginFromSquare(origin uint8) {
	*m = *m &^ (0b111111 << 26) //clear origin bits
	*m |= Move(origin) << 26    //set the cleared bits
}

//Expects a bitboard with a pop count of one, and sets the destination square of given move to that square.
func (m *Move) setDestFromBB(dest uint64) {
	*m = *m &^ (0b111111 << 20)                  //clear origin bits
	*m |= Move(bits.TrailingZeros64(dest)) << 20 //set the cleared bits
}

func (m *Move) setDestFromSquare(dest uint8) {
	*m = *m &^ (0b111111 << 20) //clear origin bits
	*m |= Move(dest) << 20      //set the cleared bits
}

func (m *Move) setMoveType(mt moveType) {
	*m = *m &^ (0b111100 >> 2) //clear bits to be safe
	*m |= Move(mt >> 2)        //set move type. We shift over 2 because moveType bits are 2 bits from the end.
}

func (m *Move) copyMoveSetType (mt moveType) Move {
	newMove := *m
	newMove.setMoveType(mt)
	return newMove
}