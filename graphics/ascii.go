package graphics

import (
	"Yerba/moveGen"
	"fmt"
)

//unicode values for pieces
const (
	BlackKing   rune = 9812
	BlackQueen  rune = 9813
	BlackRook   rune = 9814
	BlackBishop rune = 9815
	BlackKnight rune = 9816
	BlackPawn   rune = 9817

	WhiteKing   rune = 9818
	WhiteQueen  rune = 9819
	WhiteRook   rune = 9820
	WhiteBishop rune = 9821
	WhiteKnight rune = 9822
	WhitePawn   rune = 9823

	EmptySquare rune = 32
)

//PrintBoard prints out an ascii representation of the given board
func PrintBoard(b moveGen.Board) {
	boardString := [8][8]rune{}

	for i := 0; i < 64; i++ {
		boardString[i/8][i%8] = EmptySquare
	}

	for i := uint(0); i < 64; i++ {
		if (b.Pawns&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhitePawn
		}
		if (b.Pawns&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackPawn
		}
		if (b.Knights&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteKnight
		}
		if (b.Knights&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackKnight
		}
		if (b.Bishops&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteBishop
		}
		if (b.Bishops&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackBishop
		}
		if (b.Rooks&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteRook
		}
		if (b.Rooks&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackRook
		}
		if (b.Queens&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteQueen
		}
		if (b.Queens&b.BlackPieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = BlackQueen
		}
		if (b.Kings&b.WhitePieces>>i)&1 == 1 {
			boardString[7-(i/8)][i%8] = WhiteKing
		}
		if (b.Kings&b.BlackPieces>>i)&1 == 1 {
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
