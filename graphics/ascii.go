package graphics

import (
	"Yerba/moveGen"
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
	BlackKing   PieceRune = 9812
	BlackQueen  PieceRune = 9813
	BlackRook   PieceRune = 9814
	BlackBishop PieceRune = 9815
	BlackKnight PieceRune = 9816
	BlackPawn   PieceRune = 9817

	WhiteKing   PieceRune = 9818
	WhiteQueen  PieceRune = 9819
	WhiteRook   PieceRune = 9820
	WhiteBishop PieceRune = 9821
	WhiteKnight PieceRune = 9822
	WhitePawn   PieceRune = 9823

	EmptySquare PieceRune = 32
)

//PrintBoard prints out an ascii representation of the given board
func PrintBoard(b moveGen.Board) {
	boardString := [8][8]PieceRune{}

	for i := 0; i < 64; i++ {
		boardString[i/8][i%8] = EmptySquare
	}

	for i := uint(0); i < 64; i++ {
		if (b.Pawns&b.White>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhitePawn
		}
		if (b.Pawns&b.Black>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackPawn
		}
		if (b.Knights&b.White>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteKnight
		}
		if (b.Knights&b.Black>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackKnight
		}
		if (b.Bishops&b.White>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteBishop
		}
		if (b.Bishops&b.Black>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackBishop
		}
		if (b.Rooks&b.White>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteRook
		}
		if (b.Rooks&b.Black>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackRook
		}
		if (b.Queens&b.White>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteQueen
		}
		if (b.Queens&b.Black>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackQueen
		}
		if (b.Kings&b.White>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteKing
		}
		if (b.Kings&b.Black>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackKing
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
