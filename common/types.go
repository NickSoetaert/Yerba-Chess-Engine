package common

/*
Board represents one possible board orientation.
The a1 square is bit position 0, b2 = 1,..., g8 = 62, h8 = 63
*/
type Board struct {
	Pawns, Knights, Bishops, Rooks, Queens, Kings, White, Black uint64
}
