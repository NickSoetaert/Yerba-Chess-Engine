package moveGen

//Origin square bits:	   	0-5
//Destination square bits: 	6-11
//Special move bits:		12-14

//Special move flag bits:
//Denotes if a pawn promotion, castling, or en passant is taking place.
//000 - normal move
//001 - en passant
//010 - castle kingside
//011 - castle queenside
//100 - knight promotion
//101 - bishop promotion
//110 - rook promotion
//111 - queen promotion

//origin	dest		special move	unused
//111111	111111		111				1
type Move uint16

//Given a starting board and a move, return the resulting board.
func ApplyMove() {

}
