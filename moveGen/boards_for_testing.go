package moveGen

func SetUpWhiteCheckmateBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Rooks:       H1 | H2,
		Kings:       A1 | H8,
		Knights:     C3,
		WhitePieces: A1 | C3,
		BlackPieces: H1 | H2 | H8,
		RookDB:      r,
		BishopDB:    b,
		IsWhiteMove: true,
	}
}
func SetUpBlackCheckmateBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Rooks:       H1 | H2,
		Kings:       A1 | H8,
		Knights:     C3,
		WhitePieces: H1 | H2 | H8,
		BlackPieces: A1 | C3,
		RookDB:      r,
		BishopDB:    b,
		IsWhiteMove: false,
	}
}

func SetUpWhitePawnCaptureBoard() Board {
	return Board{
		Pawns:       D5,
		Knights:     C6 | E6,
		Kings:       A1 | A3,
		WhitePieces: A1 | D5,
		BlackPieces: A3 | C6 | E6,
		IsWhiteMove: true,
	}
}

func SetUpBlackPawnCaptureBoard() Board {
	return Board{
		Pawns:       D4,
		Knights:     C3 | E3,
		Kings:       A8 | C8,
		WhitePieces: A8 | C3 | E3,
		BlackPieces: C8 | D4,
		IsWhiteMove: false,
	}
}

func SetUpWhitePromotionBoard() Board {
	return Board{
		Pawns:       D7,
		Knights:     C8 | E8,
		Kings:       A1 | A3,
		WhitePieces: A1 | D7,
		BlackPieces: A3 | C8 | E8,
		IsWhiteMove: true,
	}
}

func SetUpBlackPromotionBoard() Board {
	return Board{
		Pawns:       D2,
		Knights:     C1 | E1,
		Kings:       A8 | C8,
		WhitePieces: C1 | E1 | C8,
		BlackPieces: A8 | D2,
		IsWhiteMove: false,
	}
}

func WhiteKingBoard() Board {
	return Board{
		Pawns:       C4 | D4 | E4 | C5 | E5 | C6 | E6,
		Kings:       D5 | H8,
		WhitePieces: D5,
		BlackPieces: H8 | C4 | D4 | E4 | C5 | E5 | C6 | E6,
		IsWhiteMove: true,
	}
}

func BlackKingBoard() Board {
	return Board{
		Pawns:       C4 | E4 | C5 | E5 | C6 | D6 | E6,
		Kings:       D5 | H8,
		WhitePieces: H8 | C4 | E4 | C5 | E5 | C6 | D6| E6,
		BlackPieces: D5,
		IsWhiteMove: false,
	}
}

func WhitePawnBoard() Board {
	return Board{
		Pawns:       C6 | D5 | E6,
		Kings:       F8 | H8,
		WhitePieces: D5 | F8,
		BlackPieces: H8 | C6 | E6,
		IsWhiteMove: true,
	}
}
func BlackPawnBoard() Board {
	return Board{
		Pawns:       C4 | D5 | E4,
		Kings:       F8 | H8,
		WhitePieces: H8 | C4 | E4,
		BlackPieces: D5 | F8,
		IsWhiteMove: false,
	}
}

func WhiteQueenBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       C5,
		Queens:      D4,
		RookDB:      r,
		BishopDB:    b,
		WhitePieces: D4 | H8,
		BlackPieces: C5 | F8,
		IsWhiteMove: true,
	}
}

func BlackQueenBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       C5,
		Queens:      D4,
		RookDB:      r,
		BishopDB:    b,
		WhitePieces: C5 | F8,
		BlackPieces: D4 | H8,
		IsWhiteMove: false,
	}
}

func WhiteRookBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       D5,
		Rooks:       D4,
		RookDB:      r,
		BishopDB:    b,
		BlackPieces: D5 | F8,
		WhitePieces: D4 | H8,
		IsWhiteMove: true,
	}
}
func BlackRookBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       D5,
		Rooks:       D4,
		RookDB:      r,
		BishopDB:    b,
		WhitePieces: D5 | F8,
		BlackPieces: D4 | H8,
		IsWhiteMove: false,
	}
}

func WhiteBishopBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       C5,
		Bishops:     D4,
		RookDB:      r,
		BishopDB:    b,
		BlackPieces: C5 | F8,
		WhitePieces: D4 | H8,
		IsWhiteMove: true,
	}
}
func BlackBishopBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       C5,
		Bishops:     D4,
		RookDB:      r,
		BishopDB:    b,
		WhitePieces: C5 | F8,
		BlackPieces: D4 | H8,
		IsWhiteMove: false,
	}
}

func WhiteKnightBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       C6,
		Knights:     D4,
		RookDB:      r,
		BishopDB:    b,
		BlackPieces: C6 | F8,
		WhitePieces: D4 | H8,
		IsWhiteMove: true,
	}
}

func BlackKnightBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Kings:       F8 | H8,
		Pawns:       C6,
		Knights:     D4,
		RookDB:      r,
		BishopDB:    b,
		WhitePieces: C6 | F8,
		BlackPieces: D4 | H8,
		IsWhiteMove: false,
	}
}

func SetUpBlackTakesTowardsAFileEPBoard() Board {
	return Board{
		Pawns:         A2 | B4,
		Kings:         F8 | H8,
		WhitePieces:   A2 | F8,
		BlackPieces:   B4 | H8,
		IsWhiteMove:   true,
		EnPassantFile: 0,
	}
}

func SetUpBlackTakesTowardsHFileEPBoard() Board {
	return Board{
		Pawns:         H2 | G4,
		Kings:         A1 | C1,
		WhitePieces:   H2 | A1,
		BlackPieces:   G4 | C1,
		IsWhiteMove:   true,
		EnPassantFile: 0,
	}
}

func SetUpWhiteTakesTowardsAFileEPBoard() Board {
	return Board{
		Pawns:         A7 | B5,
		Kings:         F1 | H1,
		WhitePieces:   B5 | H1,
		BlackPieces:   A7 | F1,
		IsWhiteMove:   false,
		EnPassantFile: 0,
	}
}

func SetUpWhiteTakesTowardsHFileEPBoard() Board {
	return Board{
		Pawns:         H7 | G5,
		Kings:         F1 | H1,
		WhitePieces:   G5 | H1,
		BlackPieces:   H7 | F1,
		IsWhiteMove:   false,
		EnPassantFile: 0,
	}
}

//White is in check from a pawn that can be taken E.P.
func WhiteTakesEpInCheck() Board {
	r, b := InitSlidingPieces()
	return Board{
		Pawns:         C5 | D5,
		Kings:         D4 | H8,
		WhitePieces:   D4 | D5,
		BlackPieces:   C5 | H8,
		IsWhiteMove:   true,
		EnPassantFile: 3,
		RookDB:        r,
		BishopDB:      b,
	}
}

func SetUpTestBoard() Board {
	r, b := InitSlidingPieces()
	board := Board{
		Knights:                B1 | G1 | B8 | G8,
		Bishops:                C1 | F1 | C8 | F8,
		Rooks:                  A1 | H1 | A8 | H8,
		Queens:                 D1 | D8,
		Kings:                  E1 | E8,
		WhitePieces:            FirstRank,
		BlackPieces:            EighthRank,
		RookDB:                 r,
		BishopDB:               b,
		IsWhiteMove:            true,
		WhiteKingHasNeverMoved: true,
		A1RookHasNeverMoved:    true,
		A8RookHasNeverMoved:    true,
		BlackKingHasNeverMoved: true,
		H1RookHasNeverMoved:    true,
		H8RookHasNeverMoved:    true,
		EnPassantFile:          uint8(0),
	}
	return board
}

func SetUpBoardAllPawns() Board {
	return Board{
		Pawns:                  SecondRank | SeventhRank,
		Kings:                  E1 | E8,
		WhitePieces:            E1 | SecondRank,
		BlackPieces:            E8 | SeventhRank,
		IsWhiteMove:            true,
		WhiteKingHasNeverMoved: true,
		BlackKingHasNeverMoved: true,
	}
}

func KingOnlyTestBoard() Board {
	return Board{
		Kings:                  E1 | E8,
		WhitePieces:            E1,
		BlackPieces:            E8,
		IsWhiteMove:            true,
		WhiteKingHasNeverMoved: true,
		BlackKingHasNeverMoved: true,
	}
}

func WhiteKnightOnlyBoard() Board {
	return Board{
		Kings:       E1 | E8,
		Knights:     D4,
		WhitePieces: D4 | E1,
		BlackPieces: E8,
		IsWhiteMove: true,
	}
}

func BlackKnightOnlyBoard() Board {
	return Board{
		Kings:       E1 | E8,
		Knights:     D4,
		WhitePieces: E1,
		BlackPieces: E8 | D4,
		IsWhiteMove: true,
	}
}

func SetUpWhiteCastlingBoard() Board {
	r, b := InitSlidingPieces()
	board := Board{
		Pawns: SecondRank | SeventhRank,
		//Knights:                B1 | G1 | B8 | G8,
		//Bishops:                C1 | F1 | C8 | F8,
		Rooks: A1 | H1 | A8 | H8,
		//Queens:                 D1 | D8,
		Kings:                  E1 | E8,
		WhitePieces:            A1 | E1 | H1 | SecondRank,
		BlackPieces:            A8 | E8 | H8 | SeventhRank,
		RookDB:                 r,
		BishopDB:               b,
		IsWhiteMove:            true,
		WhiteKingHasNeverMoved: true,
		A1RookHasNeverMoved:    true,
		A8RookHasNeverMoved:    true,
		BlackKingHasNeverMoved: true,
		H1RookHasNeverMoved:    true,
		H8RookHasNeverMoved:    true,
		EnPassantFile:          uint8(0),
	}
	return board
}

func SetUpBlackCastlingBoard() Board {
	r, b := InitSlidingPieces()
	board := Board{
		Pawns: SecondRank | SeventhRank,
		//Knights:                B1 | G1 | B8 | G8,
		//Bishops:                C1 | F1 | C8 | F8,
		Rooks: A1 | H1 | A8 | H8,
		//Queens:                 D1 | D8,
		Kings:                  E1 | E8,
		WhitePieces:            A1 | E1 | H1 | SecondRank,
		BlackPieces:            A8 | E8 | H8 | SeventhRank,
		RookDB:                 r,
		BishopDB:               b,
		IsWhiteMove:            false,
		WhiteKingHasNeverMoved: true,
		A1RookHasNeverMoved:    true,
		A8RookHasNeverMoved:    true,
		BlackKingHasNeverMoved: true,
		H1RookHasNeverMoved:    true,
		H8RookHasNeverMoved:    true,
		EnPassantFile:          uint8(0),
	}
	return board
}

