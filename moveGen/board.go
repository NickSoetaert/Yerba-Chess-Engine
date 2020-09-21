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

func (b Board) GenerateLegalMoves() (moves []Move) {

	attackedSquares := b.GetSquaresAttackedByOpponent()

	for _, move := range b.getPawnMoves() {
		moves = append(moves, move)
	}
	for _, move := range b.getCastlingMoves(attackedSquares) {
		moves = append(moves, move)
	}

	if b.IsWhiteMove {
		for _, move := range b.getSliderMoves(b.Bishops, true, whiteBishop) {
			moves = append(moves, move) //bishops
		}
		for _, move := range b.getSliderMoves(b.Rooks, false, whiteRook) {
			moves = append(moves, move) //rooks
		}
		for _, move := range b.getSliderMoves(b.Queens, true, whiteQueen) {
			moves = append(moves, move) //queen bishop moves
		}
		for _, move := range b.getSliderMoves(b.Queens, false, whiteQueen) {
			moves = append(moves, move) //queen rook moves
		}
		for _, move := range b.getKnightMoves() {
			moves = append(moves, move)
		}
		for _, move := range b.getNormalKingMoves(attackedSquares) {
			moves = append(moves, move)
		}
	} else { //Black moves
		for _, move := range b.getSliderMoves(b.Bishops, true, blackBishop) {
			moves = append(moves, move) //bishops
		}
		for _, move := range b.getSliderMoves(b.Rooks, false, blackRook) {
			moves = append(moves, move) //rooks
		}
		for _, move := range b.getSliderMoves(b.Queens, true, blackQueen) {
			moves = append(moves, move) //queen bishop moves
		}
		for _, move := range b.getSliderMoves(b.Queens, false, blackQueen){
			moves = append(moves, move) //queen rook moves
		}
		for _, move := range b.getKnightMoves() {
			moves = append(moves, move)
		}
		for _, move := range b.getNormalKingMoves(attackedSquares) {
			moves = append(moves, move)
		}
	}
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
		fmt.Println("-------------------------------------------")
		for j := 0; j < 8; j++ {
			if j == 0 {
				fmt.Printf("%c %v", rune(65+(7-i)), "| ")
			} else {
				fmt.Printf("|  ")
			}
			fmt.Printf("%c", boardString[i][j])
			fmt.Printf(" ")
			if j == 7 {
				fmt.Printf("|")
			}
		}
		fmt.Println("")
	}
	fmt.Println("-------------------------------------------")

	fmt.Println("  | A | B  | C  | D  | E  | F  | G  | H  |")
}
