package main

import (
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

//Sets a single rook on the given square
func RookDebugBoard(squares uint64) Board {

	board := Board{
		Rooks:       squares,
		White:       squares,
		IsWhiteMove: true,
	}
	return board
}

//PrintBoard prints out an ascii representation of the given board
func (b Board) PrintBoard() {
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
