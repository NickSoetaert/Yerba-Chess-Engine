package moveGen

import "fmt"

func (b *Board) getAttackedSquareOccupancy(m Move) tileOccupancy {
	tile := m.getDestSquare()

	if b.IsWhiteMove {
		if tile & b.BlackPieces == EmptyBoard {
			return empty
		} else if tile & b.Pawns != EmptyBoard {
			return blackPawn
		} else if tile & b.Knights != EmptyBoard {
			return blackKnight
		} else if tile & b.Bishops != EmptyBoard {
			return blackBishop
		} else if tile & b.Queens != EmptyBoard {
			return blackQueen
		} else if tile & b.Kings != EmptyBoard { //todo: take this out once I add checkmate
			fmt.Println("capturing a king!")
			return blackKing
		} else {
			panic("impossible tile occupancy")
		}
	} else {
		if tile & b.BlackPieces == EmptyBoard {
			return empty
		} else if tile & b.Pawns != EmptyBoard {
			return whitePawn
		} else if tile & b.Knights != EmptyBoard {
			return whiteKnight
		} else if tile & b.Bishops != EmptyBoard {
			return whiteBishop
		} else if tile & b.Queens != EmptyBoard {
			return whiteQueen
		} else if tile & b.Kings != EmptyBoard { //todo: take this out once I add checkmate
			fmt.Println("capturing a king!")
			return whiteKing
		} else {
			panic("impossible tile occupancy")
		}
	}
}
