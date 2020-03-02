package main

/*
Board represents one possible board orientation.
The a1 square is bit position 0, b2 = 1,..., g8 = 62, h8 = 63
To figure out if , AND together with White or Black
*/
type Board struct {
	Pawns, Knights, Bishops, Rooks, Queens, Kings, White, Black uint64
	RookDB, BishopDB                                            [][]uint64
	IsWhiteMove                                                 bool
}

//SetUpBoard inits a board in the default state
func SetUpBoard() Board {
	r, b := InitSlidingPieces()
	board := Board{
		Pawns:       SecondRank | SeventhRank,
		Knights:     B1 | G1 | B8 | G8,
		Bishops:     C1 | F1 | C8 | F8,
		Rooks:       A1 | H1 | A8 | H8,
		Queens:      D1 | D8,
		Kings:       E1 | E8,
		White:       FirstRank | SecondRank,
		Black:       SeventhRank | EighthRank,
		RookDB:      r,
		BishopDB:    b,
		IsWhiteMove: true,
	}
	return board
}

func (b Board) GenerateLegalMoves() []Board {



	return nil
}

