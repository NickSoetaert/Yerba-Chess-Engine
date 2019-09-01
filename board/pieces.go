package board

import (
	//"Yerba/utils"
	"fmt"
)

//This file should declare the logic of how individual pieces move

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

func PrintBoard(board Board) {
	boardString := [8][8]PieceRune{}

	for i := 0; i < 64; i++ {
		boardString[i/8][i%8] = 32 //32 is a space in ascii
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
		for j := 0; j < 8; j++ {
			fmt.Printf("%c", boardString[i][j])
		}
		fmt.Println("")
	}
}
