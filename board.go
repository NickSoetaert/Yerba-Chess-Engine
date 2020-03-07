package main

import "Yerba/moveGen"

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
	r, b := moveGen.InitSlidingPieces()
	board := Board{
		Pawns:       moveGen.SecondRank | moveGen.SeventhRank,
		Knights:     moveGen.B1 | moveGen.G1 | moveGen.B8 | moveGen.G8,
		Bishops:     moveGen.C1 | moveGen.F1 | moveGen.C8 | moveGen.F8,
		Rooks:       moveGen.A1 | moveGen.H1 | moveGen.A8 | moveGen.H8,
		Queens:      moveGen.D1 | moveGen.D8,
		Kings:       moveGen.E1 | moveGen.E8,
		White:       moveGen.FirstRank | moveGen.SecondRank,
		Black:       moveGen.SeventhRank | moveGen.EighthRank,
		RookDB:      r,
		BishopDB:    b,
		IsWhiteMove: true,
	}
	return board
}

func (b Board) GenerateLegalMoves() []Board {

	return nil
}
