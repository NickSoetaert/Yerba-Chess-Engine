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
	//CurrentKingInCheck							  bool 		 //True if the color who's turn it is starts their turn in check

	WhiteKingHasNeverMoved bool //True if the white king has never moved (including castling)
	A1RookHasNeverMoved    bool //True if white rook on A1 has never moved, regardless of if captured or not.
	A8RookHasNeverMoved    bool //True if white rook on A8 has never moved, regardless of if captured or not.

	BlackKingHasNeverMoved bool //True if the black king has never moved (including castling)
	H1RookHasNeverMoved    bool //True if black rook on H1 has never moved, regardless of if captured or not.
	H8RookHasNeverMoved    bool //True if black rook on H8 has never moved, regardless of if captured or not.

	EnPassantFile uint8 //File on which an E.P. capture is currently legal A=1, H=8. 0 if no E.P. is legal.
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

func BlackPawnBoard() Board {
	return Board{
		Pawns:       A5 | D5 | H6,
		Kings:       F8 | H8,
		WhitePieces: H8,
		BlackPieces: A5 | D5 | H6 | F8,
		IsWhiteMove: false,
	}
}
func WhitePawnBoard() Board {
	return Board{
		Pawns:       A5 | D5 | H6,
		Kings:       F8 | H8,
		WhitePieces: A5 | D5 | H6 | F8,
		BlackPieces: H8,
		IsWhiteMove: true,
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

//For benchmarking and testing
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

// Todo: keep part of this cached from previous move.
func (b Board) GetSquaresAttackedThisHalfTurn() (defendedSquares uint64) {
	//Pretend it is the opponent's move to see what squares they can currently attack
	defendedSquares |= b.GetPawnDefendedSquares()
	defendedSquares |= b.getKnightDefendedSquares()

	defendedSquares |= b.getSliderDefendedSquares(b.Bishops, true) //bishops
	defendedSquares |= b.getSliderDefendedSquares(b.Rooks, false)  //rooks
	defendedSquares |= b.getSliderDefendedSquares(b.Queens, true)  //queen bishop moves
	defendedSquares |= b.getSliderDefendedSquares(b.Queens, false) //queen rook moves

	defendedSquares |= b.getKingDefendedSquares()

	return defendedSquares
}

// Todo: keep part of this cached from previous move.
func (b Board) GetSquaresAttackedByOpponent() (defendedSquares uint64) {
	//Pretend it is the opponent's move to see what squares they can currently attack
	b.IsWhiteMove = !b.IsWhiteMove
	return b.GetSquaresAttackedThisHalfTurn()
}

func currentTurnKingIsInCheck(king uint64, attackedSquares uint64) bool {
	return king&attackedSquares != 0
}

//Todo: account for pinned pieces
//Todo: account for check/checkmate
func (b Board) GenerateLegalMoves(pChan, nChan, bChan, rChan, qbChan, qrChan, kChan, castleChan chan[]Move) (moves []Move) {

	attackedSquares := b.GetSquaresAttackedByOpponent()

	go b.getPawnMoves(pChan) //pawns
	go b.getCastlingMoves(attackedSquares, castleChan)

	if b.IsWhiteMove {
		go b.getSliderMoves(b.Bishops, true, bChan, whiteBishop) //bishops
		go b.getSliderMoves(b.Rooks, false, rChan, whiteRook)    //rooks
		go b.getSliderMoves(b.Queens, true, qbChan, whiteQueen)  //queens
		go b.getSliderMoves(b.Queens, false, qrChan, whiteQueen) //queens

		go b.getKnightMoves(nChan)
		go b.getNormalKingMoves(attackedSquares, kChan)

	} else {
		go b.getSliderMoves(b.Bishops, true, bChan, blackBishop) //bishops
		go b.getSliderMoves(b.Rooks, false, rChan, blackRook)    //rooks
		go b.getSliderMoves(b.Queens, true, qbChan, blackQueen)  //queens
		go b.getSliderMoves(b.Queens, false, qrChan, blackQueen) //queens

		go b.getKnightMoves(nChan)
		go b.getNormalKingMoves(attackedSquares, kChan)
	}

	moves = append(moves, <-pChan...)
	moves = append(moves, <-nChan...)
	moves = append(moves, <-bChan...)
	moves = append(moves, <-rChan...)
	moves = append(moves, <-qbChan...)
	moves = append(moves, <-qrChan...)
	moves = append(moves, <-kChan...)
	moves = append(moves, <-castleChan...)

	return moves
}

//PrintBoard prints out an ascii representation of the given board
func PrintBoard(b Board) {
	boardString := [8][8]rune{}

	for i := 0; i < 64; i++ {
		boardString[i/8][i%8] = EmptySquareIcon
	}

	for i := uint(0); i < 64; i++ {
		if (b.Pawns&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhitePawnIcon
		}
		if (b.Pawns&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackPawnIcon
		}
		if (b.Knights&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteKnightIcon
		}
		if (b.Knights&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackKnightIcon
		}
		if (b.Bishops&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteBishopIcon
		}
		if (b.Bishops&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackBishopIcon
		}
		if (b.Rooks&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteRookIcon
		}
		if (b.Rooks&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackRookIcon
		}
		if (b.Queens&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteQueenIcon
		}
		if (b.Queens&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackQueenIcon
		}
		if (b.Kings&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteKingIcon
		}
		if (b.Kings&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackKingIcon
		}
	}
	for i := 0; i < 8; i++ {
		fmt.Println("-----------------------------------------")
		for j := 0; j < 8; j++ {

			fmt.Printf("|  ")
			fmt.Printf("%c", boardString[i][j])
			fmt.Printf(" ")
			if j == 7 {
				fmt.Printf("|")
			}
		}
		fmt.Println("")
	}
	fmt.Println("-----------------------------------------")
}
