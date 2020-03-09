package moveGen

//type Move structure:
//Origin square bits:	   	0-5
//Destination square bits: 	6-11
//Special move bits:		12-14
//Who's turn:				15

//Origin/dest square bits:
//0 represents A1, and 64 represents H8

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

//Turn bits:
//0 - white to move
//1 - black to move

//000000 000001 000 1 : A white piece started on A1 and moved to A2.
//000000 000000 010 1 : white castled kingside
//010011 010100 001 0 : black captured en passant from d4 to e3

type Move uint16

type undoMove func()


//TODO
//Given a starting board and a move, return the resulting board.
//Returns a function that undoes the previously applied move.
func (b *Board) ApplyMove(m Move) undoMove {

	return func(){
		b.White = 0
	}
}
