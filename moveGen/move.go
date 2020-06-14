package moveGen

import (
	"math/bits"
)

//type Move:
//Origin square bits:	   	0-5
//Destination square bits: 	6-11
//Special move bits:		12-15

//Origin/dest square bits:
//0 represents A1, and 64 represents H8

//Special move flag bits
//Note that all flags are mutually exclusive
//0000 - normal pawn push
//0001 - normal pawn capture
//0010 - double pawn push
//0011 - en passant capture
//0100 - knight move
//0101 - bishop move
//0110 - rook move
//0111 - queen move
//1000 - king move
//1001 - castle kingside
//1010 - castle queenside
//1011 - knight promotion
//1100 - bishop promotion
//1101 - rook promotion
//1110 - queen promotion
//1111 - UNUSED

type Move uint16

type UndoMove func()

//Given a starting board and a move, return the resulting board.
//Returns a function that undoes the previously applied move.
//ApplyMove does NOT check for legality; that is the responsibility of MoveGen.
func (b *Board) ApplyMove(m Move) UndoMove {
	oldBoard := *b //TODO: Optimize

	//Note - adding/removing pieces from the White/Black bitboards is taken care of in a single case after the switch.
	switch m.getMoveType() {
	case normalPawnPush:
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.Pawns |= m.getDest()             //add pawn

	case normalPawnCapture:
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.clearTargetSquare(m.getDest())   //clear captured piece
		b.Pawns |= m.getDest()             //add pawn

	case pawnDoublePush:
		b.EnPassantFile = uint8(bits.TrailingZeros64(m.getOrigin()) % 8)
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.Pawns |= m.getDest()             //add pawn

	case enPassantCapture:
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.Pawns |= m.getDest()             //add pawn
		if b.IsWhiteMove {
			b.clearTargetSquare(uint64(b.EnPassantFile) << 16) //remove captured piece
		} else {
			b.clearTargetSquare(uint64(b.EnPassantFile) << 40) //TODO: check
		}

	case knightMove:
		b.Knights = b.Knights &^ m.getOrigin()
		b.clearTargetSquare(m.getDest())
		b.Knights |= m.getDest()

	case bishopMove:
		b.Bishops = b.Bishops &^ m.getOrigin()
		b.clearTargetSquare(m.getDest())
		b.Bishops |= m.getDest()

	case rookMove:
		//It's possible that we were moving a non-original rook from a corner square, but that doesn't matter as
		//the original rook is no longer on that corner square. (So castle rights were revoked in the first place.)
		if m.getOrigin() == A1 {
			b.WhiteQueensideCastleRights = false
		} else if m.getOrigin() == H1 {
			b.WhiteKingsideCastleRights = false
		} else if m.getOrigin() == A8 {
			b.BlackQueensideCastleRights = false
		} else if m.getOrigin() == H8 {
			b.BlackKingsideCastleRights = false
		}

		b.Rooks = b.Rooks &^ m.getOrigin()
		b.clearTargetSquare(m.getDest())
		b.Rooks |= m.getDest()

	case queenMove:
		b.Queens = b.Queens &^ m.getOrigin()
		b.clearTargetSquare(m.getDest())
		b.Queens |= m.getDest()

	case normalKingMove:
		if b.IsWhiteMove {
			b.WhiteKingsideCastleRights = false
			b.WhiteQueensideCastleRights = false
		} else {
			b.BlackKingsideCastleRights = false
			b.WhiteQueensideCastleRights = false
		}
		b.Kings = b.Kings &^ m.getOrigin()
		b.clearTargetSquare(m.getDest())
		b.Kings |= m.getDest()

	case castleKingside:
		b.Kings = b.Kings &^ m.getOrigin()
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
		b.Kings = b.Kings &^ m.getOrigin()
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
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.clearTargetSquare(m.getDest())   //clear possible capture
		b.Knights |= m.getDest()           //add knight

	case bishopPromotion:
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.clearTargetSquare(m.getDest())   //clear possible capture
		b.Bishops |= m.getDest()           //add bishop

	case rookPromotion:
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.clearTargetSquare(m.getDest())   //clear possible capture
		b.Rooks |= m.getDest()             //add rook

	case queenPromotion:
		b.Pawns = b.Pawns &^ m.getOrigin() //clear pawn
		b.clearTargetSquare(m.getDest())   //clear possible capture
		b.Queens |= m.getDest()            //add queen

	}

	//Clear en passant capture unless we just pushed a pawn two squares
	if m.getMoveType() != pawnDoublePush {
		b.EnPassantFile = 0
	}

	//Take care of adding/removing pieces from the color bitboards.
	if b.IsWhiteMove {
		b.WhitePieces = b.WhitePieces &^ m.getOrigin()
		b.WhitePieces |= m.getDest()
	} else {
		b.BlackPieces = b.BlackPieces &^ m.getOrigin()
		b.BlackPieces |= m.getDest()
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
func (m Move) getOrigin() uint64 {
	return 1 << (m >> 10)
}

//001000 010000 0000
//		 100000
//1 00000000 00000000
func (m Move) getDest() uint64 {
	m &= 0b0000001111110000
	return 1 << (m >> 4)
}

func (m Move) getMoveType() moveType {
	return moveType(m & 0b0000000000001111)
}

//Expects a bitboard with a pop count of one, and sets the origin square of given move to that square.
func (m *Move) setOriginFromBB(origin uint64) {
	*m = *m &^ (0b111111 << 10)                    //clear origin bits
	*m |= Move(bits.TrailingZeros64(origin)) << 10 //set the cleared bits
}

//Expects a square position
func (m *Move) setOriginFromSquare(origin uint8) {
	*m = *m &^ (0b111111 << 10) //clear origin bits
	*m |= Move(origin) << 10    //set the cleared bits
}

//Expects a bitboard with a pop count of one, and sets the destination square of given move to that square.
func (m *Move) setDestFromBB(dest uint64) {
	*m = *m &^ (0b111111 << 4)                  //clear origin bits
	*m |= Move(bits.TrailingZeros64(dest)) << 4 //set the cleared bits
}

func (m *Move) setDestFromSquare(dest uint8) {
	*m = *m &^ (0b111111 << 4) //clear origin bits
	*m |= Move(dest) << 4      //set the cleared bits
}

func (m *Move) setMoveType(mt moveType) {
	*m = *m &^ 0b111111 //clear origin bits
	*m |= Move(mt)      //set the cleared bits
}

func (m *Move) copyMoveSetType (mt moveType) Move {
	newMove := *m
	newMove.setMoveType(mt)
	return newMove
}