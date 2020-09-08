package moveGen

import (
	"Yerba/utils"
	"fmt"
)

func (b *Board) getTileOccupancy(tile uint64) tileOccupancy {

	if b.IsWhiteMove {
		if tile&b.BlackPieces == EmptyBoard {
			return empty
		} else if tile&b.Pawns != EmptyBoard {
			return blackPawn
		} else if tile&b.Knights != EmptyBoard {
			return blackKnight
		} else if tile&b.Bishops != EmptyBoard {
			return blackBishop
		} else if tile&b.Rooks != EmptyBoard {
			return blackRook
		} else if tile&b.Queens != EmptyBoard {
			return blackQueen
		} else if tile&b.Kings != EmptyBoard { //todo: take this out once I add checkmate
			fmt.Println("capturing a king!")
			return blackKing
		} else {
			PrintBoard(*b)
			utils.PrintBinaryBoard(tile)
			fmt.Printf("tile and white == empty: %v\n", tile&b.WhitePieces == EmptyBoard)
			fmt.Printf("tile and black == empty: %v\n", tile&b.BlackPieces == EmptyBoard)
			fmt.Printf("tile and pawn == empty: %v\n", tile&b.Pawns == EmptyBoard)
			fmt.Printf("tile and kn == empty: %v\n", tile&b.Knights == EmptyBoard)
			fmt.Printf("tile and b == empty: %v\n", tile&b.Bishops == EmptyBoard)
			fmt.Printf("tile and r == empty: %v\n", tile&b.Rooks == EmptyBoard)
			fmt.Printf("tile and q == empty: %v\n", tile&b.Queens == EmptyBoard)
			fmt.Printf("tile and k == empty: %v\n", tile&b.Kings == EmptyBoard)

			panic("white move - impossible tile occupancy")
		}
	} else {
		if tile&b.BlackPieces == EmptyBoard {
			return empty
		} else if tile&b.Pawns != EmptyBoard {
			return whitePawn
		} else if tile&b.Knights != EmptyBoard {
			return whiteKnight
		} else if tile&b.Bishops != EmptyBoard {
			return whiteBishop
		} else if tile&b.Rooks != EmptyBoard {
			return whiteRook
		} else if tile&b.Queens != EmptyBoard {
			return whiteQueen
		} else if tile&b.Kings != EmptyBoard { //todo: take this out once I add checkmate
			fmt.Println("capturing a king!")
			return whiteKing
		} else {
			PrintBoard(*b)
			utils.PrintBinaryBoard(tile)
			fmt.Printf("tile and white == empty: %v\n", tile&b.WhitePieces == EmptyBoard)
			fmt.Printf("tile and black == empty: %v\n", tile&b.BlackPieces == EmptyBoard)
			fmt.Printf("tile and pawn == empty: %v\n", tile&b.Pawns == EmptyBoard)
			fmt.Printf("tile and kn == empty: %v\n", tile&b.Knights == EmptyBoard)
			fmt.Printf("tile and b == empty: %v\n", tile&b.Bishops == EmptyBoard)
			fmt.Printf("tile and r == empty: %v\n", tile&b.Rooks == EmptyBoard)
			fmt.Printf("tile and q == empty: %v\n", tile&b.Queens == EmptyBoard)
			fmt.Printf("tile and k == empty: %v\n", tile&b.Kings == EmptyBoard)

			panic("black move - impossible tile occupancy")
		}
	}
}
