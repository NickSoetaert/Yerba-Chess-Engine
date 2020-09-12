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
func SetUpCheckmateBoard() Board {
	r, b := InitSlidingPieces()
	return Board{
		Rooks:                  H1 | H2 | C6,
		Kings:                  A1 | H8,
		WhitePieces:            A1 | C6,
		BlackPieces:            H1 | H2 | H8,
		RookDB:                 r,
		BishopDB:               b,
		IsWhiteMove:            true,
	}
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
		Pawns: SecondRank | SeventhRank,
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

func KnightOnlyTestBoard() Board {
	return Board{
		Kings:                  E1 | E8,
		Knights: B1 | G1 | B8 | G8,
		WhitePieces:            B1 | G1 | E1,
		BlackPieces:            B8 | G8 | E8,
		IsWhiteMove:            true,
	}
}

//For testing
func SetUpCastlingTestBoard() Board {
	r, b := InitSlidingPieces()
	board := Board{
		Kings:                  E1 | E8,
		WhitePieces:            E1 | A1 | H1,
		BlackPieces:            E8,
		Rooks:                  A1 | H1,
		RookDB:                 r,
		BishopDB:               b,
		IsWhiteMove:            true,
		WhiteKingHasNeverMoved: true,
		H8RookHasNeverMoved:    true,
		A1RookHasNeverMoved:    true,
		BlackKingHasNeverMoved: true,
		H1RookHasNeverMoved:    true,
		A8RookHasNeverMoved:    true,
		EnPassantFile:          uint8(0),
	}
	return board
}

// Todo: keep part of this cached from previous move.
func (b Board) GetSquaresAttackedThisHalfTurn() (defendedSquares uint64) {
	//Pretend it is the opponent's move to see what squares they can currently attack
	defendedSquares |= b.getPawnDefendedSquares()
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
	b.IsWhiteMove = ! b.IsWhiteMove

	defendedSquares |= b.getPawnDefendedSquares()
	defendedSquares |= b.getKnightDefendedSquares()

	defendedSquares |= b.getSliderDefendedSquares(b.Bishops, true) //bishops
	defendedSquares |= b.getSliderDefendedSquares(b.Rooks, false)  //rooks
	defendedSquares |= b.getSliderDefendedSquares(b.Queens, true)  //queen bishop moves
	defendedSquares |= b.getSliderDefendedSquares(b.Queens, false) //queen rook moves

	defendedSquares |= b.getKingDefendedSquares()


	return defendedSquares
}

func currentTurnKingIsInCheck(king uint64, attackedSquares uint64) bool {
	return king & attackedSquares != 0
}

//Todo: account for pinned pieces
//Todo: account for check/checkmate
func (b Board) GenerateLegalMoves() (moves []Move) {
	pChan := make(chan []Move)
	nChan := make(chan []Move)
	bChan := make(chan []Move)
	rChan := make(chan []Move)
	qbChan := make(chan []Move)
	qrChan := make(chan []Move)
	kChan := make(chan []Move)
	castleChan := make(chan []Move)

	attackedSquares := b.GetSquaresAttackedByOpponent()


	if b.IsWhiteMove {
		if currentTurnKingIsInCheck(b.Kings & b.WhitePieces, attackedSquares) {
			fmt.Println("White king in check!")
			PrintBoard(b)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~")
		}
		go b.getSliderMoves(b.Bishops, true, bChan, whiteBishop) //bishops
		go b.getSliderMoves(b.Rooks, false, rChan, whiteRook)    //rooks
		go b.getSliderMoves(b.Queens, true, qbChan, whiteQueen)  //queens
		go b.getSliderMoves(b.Queens, false, qrChan, whiteQueen) //queens

		go b.getKnightMoves(nChan)
		go b.getNormalKingMoves(attackedSquares, kChan)

	} else {

		if currentTurnKingIsInCheck(b.Kings & b.BlackPieces, attackedSquares) {
			fmt.Println("Black King in check!")
			PrintBoard(b)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~")
		}
		go b.getSliderMoves(b.Bishops, true, bChan, blackBishop) //bishops
		go b.getSliderMoves(b.Rooks, false, rChan, blackRook)    //rooks
		go b.getSliderMoves(b.Queens, true, qbChan, blackQueen)  //queens
		go b.getSliderMoves(b.Queens, false, qrChan, blackQueen) //queens

		go b.getKnightMoves(nChan)
		go b.getNormalKingMoves(attackedSquares, kChan)
	}

	go b.getPawnMoves(pChan) //pawns
	go b.getCastlingMoves(attackedSquares, castleChan)

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
