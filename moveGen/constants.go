package moveGen

//type Move: 32 bit unsigned int.
//Origin square coordinate bits: 0-5 (0 represents A1, and 64 represents H8)
//Destination square coordinate bits: 6-11
//Origin square occupancy type: 12-15
//Destination square occupancy type PRE MOVE: 16-20
//Destination square occupancy POST MOVE: 21-25
//Move type bits (needed for undo move function): 26-29
//Who's move: bit 30. (0 means White made the move, 1 means Black made the move.)
//bit 31: currently unused
type Move uint32

const (
	originSquareBitsStart = 0
	originSquareBitsEnd = 5
	destSquareBitsStart = 6
	destSquareBitsEnd = 11
	originSquareOccBitsStart = 12
	originSquareOccBitsEnd = 15
	destSquarePreMoveOccBitsStart = 16
	destSquarePreMoveOccBitsEnd = 20
	destSquarePostMoveOccBitsStart = 21
	destSquarePostMoveOccBitsEnd = 25
	moveTypeBitsStart = 26
	moveTypeBitsEnd = 28
	whoseTurnBit = 29
)

//"move type" bits. These are the 4 bits inside a Move at bits 26-29
//000 - normal move
//001 - double pawn push (Needed for en passant)
//010 - en passant capture
//011 - castle kingside
//100 - castle queenside
type moveType uint8

const (
	normalMove moveType = iota
	pawnDoublePush
	enPassantCapture
	castleKingside
	castleQueenside
)

//tile occupancy types - the piece (or lack of) that's on a square.
type tileOccupancy uint8

const (
	empty tileOccupancy = iota

	whitePawn
	whiteKnight
	whiteBishop
	whiteRook
	whiteQueen
	whiteKing

	blackPawn
	blackKnight
	blackBishop
	blackRook
	blackQueen
	blackKing
)
