package moveGen

type moveType uint8

const (
	normalMove moveType = iota
	pawnDoublePush
	enPassantCapture
	removesWhiteKingsideCastleRights
	removesWhiteQueensideCastleRights
	removesBlackKingsideCastleRights
	removesBlackQueensideCastleRights
	castleKingside
	castleQueenside
)

//tile occupancy types
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
