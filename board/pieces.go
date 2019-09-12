package board

import (
	//"Yerba/utils"
	"fmt"
)

//PieceType - pawn, knight, etc.
type PieceType uint8

const (
	Pawn PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)

//PieceRune is an unicode value for ascii piece representation
type PieceRune rune

//unicode values for pieces
const (
	WhiteKing   PieceRune = 9812
	WhiteQueen  PieceRune = 9813
	WhiteRook   PieceRune = 9814
	WhiteBishop PieceRune = 9815
	WhiteKnight PieceRune = 9816
	WhitePawn   PieceRune = 9817

	BlackKing   PieceRune = 9818
	BlackQueen  PieceRune = 9819
	BlackRook   PieceRune = 9820
	BlackBishop PieceRune = 9821
	BlackKnight PieceRune = 9822
	BlackPawn   PieceRune = 9823

	EmptySquare PieceRune = 32
)

//IsWhite == true means the piece, square or turn is White or White's. Else, it's Black or Black's.
type IsWhite bool

//Defines if a piece, tile or turn is white or black
const (
	White IsWhite = true
	Black IsWhite = false
)

//SetUpBoard inits a board in the default state
func SetUpBoard() Board {
	board := Board{
		Pawns:   SecondRank | SeventhRank,
		Knights: B1 | G1 | B8 | G8,
		Bishops: C1 | F1 | C8 | F8,
		Rooks:   A1 | H1 | A8 | H8,
		Queens:  D1 | D8,
		Kings:   E1 | E8,
		White:   FirstRank | SecondRank,
		Black:   SeventhRank | EighthRank,
		Move:    White,
	}

	return board
}

//PrintBoard prints out an ascii representation of the given board
func PrintBoard(board Board) {
	boardString := [8][8]PieceRune{}

	for i := 0; i < 64; i++ {
		boardString[i/8][i%8] = EmptySquare
	}

	for i := uint(0); i < 64; i++ {
		if (board.Pawns&board.White>>i)&1 == 1 {
			boardString[i/8][i%8] = WhitePawn
		}
		if (board.Pawns&board.Black>>i)&1 == 1 {
			boardString[i/8][i%8] = BlackPawn
		}
		if (board.Knights&board.White>>i)&1 == 1 {
			boardString[i/8][i%8] = WhiteKnight
		}
		if (board.Knights&board.Black>>i)&1 == 1 {
			boardString[i/8][i%8] = BlackKnight
		}
		if (board.Bishops&board.White>>i)&1 == 1 {
			boardString[i/8][i%8] = WhiteBishop
		}
		if (board.Bishops&board.Black>>i)&1 == 1 {
			boardString[i/8][i%8] = BlackBishop
		}
		if (board.Rooks&board.White>>i)&1 == 1 {
			boardString[i/8][i%8] = WhiteRook
		}
		if (board.Rooks&board.Black>>i)&1 == 1 {
			boardString[i/8][i%8] = BlackRook
		}

		if (board.Queens&board.White>>i)&1 == 1 {
			boardString[i/8][i%8] = WhiteQueen
		}

		if (board.Queens&board.Black>>i)&1 == 1 {
			boardString[i/8][i%8] = BlackQueen
		}
		if (board.Kings&board.White>>i)&1 == 1 {
			boardString[i/8][i%8] = WhiteKing
		}
		if (board.Kings&board.Black>>i)&1 == 1 {
			boardString[i/8][i%8] = BlackKing
		}
	}
	for i := 0; i < 8; i++ {
		fmt.Println("---------------------------------")
		for j := 0; j < 8; j++ {

			fmt.Printf("| ")
			fmt.Printf("%c", boardString[i][j])
			fmt.Printf(" ")
			if j == 7 {
				fmt.Printf("|")
			}
		}
		fmt.Println("")
	}
	fmt.Println("---------------------------------")
}
