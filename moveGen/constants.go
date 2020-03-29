package moveGen

type moveType uint8

const (
	normalPawnPush moveType = iota
	normalPawnCapture
	pawnDoublePush
	enPassantCapture
	knightMove
	bishopMove
	rookMove
	queenMove
	normalKingMove
	castleKingside
	castleQueenside
	knightPromotion
	bishopPromotion
	rookPromotion
	queenPromotion
)