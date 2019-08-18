package board

/*
Board represents a game board state
A Board is made up of 12 unsigned 64 bit ints, each representing one piece type.
The a1 square is bit position 0, b2 = 1,..., g8 = 62, h8 = 63
*/ 
type Board struct {
	WP uint64
	WN uint64
	WB uint64
	WR uint64
	WQ uint64
	WK uint64

	BP uint64
	BN uint64
	BB uint64
	BR uint64
	BQ uint64
	BK uint64
}


