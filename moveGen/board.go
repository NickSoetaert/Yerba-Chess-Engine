package moveGen

import (
	"fmt"
)

//An instance of a Board represents a single possible game state.
//The a1 square is bit position 0, b2 = 1,..., g8 = 62, h8 = 63
//EnPassantFile encoding: Square that en passant is currently possible on. To get square, AND with turn.
//00000000 == NONE
//10000000 == A
//01000000 == B
//00100000 == C
//00010000 == D
//00001000 == E
//00000100 == F
//00000010 == G
//00000001 == H

type Board struct {
	Pawns, Knights, Bishops, Rooks, Queens, Kings uint64
	WhitePieces, BlackPieces                      uint64
	RookDB, BishopDB                              [][]uint64 //All possible (precomputed) moves for rooks and bishops
	IsWhiteMove                                   bool       //True if it is currently white's move

	WhiteKingHasNeverMoved bool //True if the white king has never moved (including castling)
	A1RookHasNeverMoved    bool //True if white rook on A1 has never moved, regardless of if captured or not.
	A8RookHasNeverMoved    bool //True if white rook on A8 has never moved, regardless of if captured or not.

	BlackKingHasNeverMoved bool //True if the black king has never moved (including castling)
	H1RookHasNeverMoved    bool //True if black rook on H1 has never moved, regardless of if captured or not.
	H8RookHasNeverMoved    bool //True if black rook on H8 has never moved, regardless of if captured or not.

	EnPassantFile uint8 //File on which an E.P. capture is currently legal, 0 if no E.P. is legal.
}

//SetUpBoard inits a board in the default state
func SetUpBoard() Board {
	r, b := InitSlidingPieces()
	board := Board{
		Pawns:                  SecondRank | SeventhRank,
		Knights:                B1 | G1 | B8 | G8,
		Bishops:                C1 | F1 | C8 | F8,
		Rooks:                  A1 | H1 | A8 | H8,
		Queens:                 D1 | D8,
		Kings:                  E1 | E8,
		WhitePieces:            FirstRank | SecondRank,
		BlackPieces:            SeventhRank | EighthRank,
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

//For benchmarking
func SetUpBoardNoPawns() Board {
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

//Todo: account for pinned pieces
func (b Board) GenerateLegalMoves() (moves []Move) {
	pChan := make(chan []Move)
	nChan := make(chan []Move)
	bChan := make(chan []Move)
	rChan := make(chan []Move)
	qbChan := make(chan []Move)
	qrChan := make(chan []Move)
	kChan := make(chan []Move)
	castleChan := make(chan []Move)

	go b.getPawnMoves(pChan)                                                                           //pawns
	go getSliderMoves(b.Bishops, b.WhitePieces, b.BlackPieces, b.IsWhiteMove, true, b.BishopDB, bChan) //bishops
	go getSliderMoves(b.Rooks, b.WhitePieces, b.BlackPieces, b.IsWhiteMove, false, b.RookDB, rChan)    //rooks
	go getSliderMoves(b.Queens, b.WhitePieces, b.BlackPieces, b.IsWhiteMove, true, b.BishopDB, qbChan) //queens
	go getSliderMoves(b.Queens, b.WhitePieces, b.BlackPieces, b.IsWhiteMove, false, b.RookDB, qrChan)  //queens
	go b.getCastlingMoves(EmptyBoard, castleChan)                                                      //Todo: pass attacked squares

	if b.IsWhiteMove {
		go getKnightMoves(b.Knights, b.WhitePieces, nChan)
		go getNormalKingMoves(b.Kings, b.WhitePieces, EmptyBoard, kChan) //Todo: pass attacked squares
	} else {
		go getKnightMoves(b.Knights, b.BlackPieces, nChan)
		go getNormalKingMoves(b.Kings, b.BlackPieces, EmptyBoard, kChan)
	}

	moves = append(moves, <-pChan...)
	fmt.Printf("number of pawn moves: %v\n", len(moves))
	moves = append(moves, <-nChan...)
	fmt.Printf("knight moves: %v\n",len(moves))
	moves = append(moves, <-bChan...)
	fmt.Printf("bishop moves: %v\n",len(moves))
	moves = append(moves, <-rChan...)
	fmt.Printf("rook moves: %v\n",len(moves))
	moves = append(moves, <-qbChan...)
	fmt.Printf("queen bishop moves: %v\n",len(moves))
	moves = append(moves, <-qrChan...)
	fmt.Printf("queen rook moves: %v\n",len(moves))
	moves = append(moves, <-kChan...)
	fmt.Printf("king moves: %v\n",len(moves))
	moves = append(moves, <-castleChan...)
	fmt.Printf("castle moves: %v\n",len(moves))

	return moves
}
