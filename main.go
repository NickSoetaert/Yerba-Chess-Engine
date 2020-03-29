//Nick Soetaert
package main

import (
	"Yerba/graphics"
	"Yerba/moveGen"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//40==A6

	b := moveGen.SetUpBoard()

	pawnMoves := moveGen.GetPawnMoves(b.Pawns, b.White, b.Black, b.IsWhiteMove, b.EnPassantFile)
	for _, move := range pawnMoves {
		undo := b.ApplyMove(move)
		graphics.PrintBoard(b)
		deeper := moveGen.GetPawnMoves(b.Pawns,  b.White, b.Black, b.IsWhiteMove, b.EnPassantFile)
		for _, blackMove := range deeper {
			blackUndo := b.ApplyMove(blackMove)
			graphics.PrintBoard(b)
			blackUndo()
		}
		fmt.Println("")
		undo()
	}

	//for _, move := range moveGen.GetPawnMoves(b.Pawns, b.White, b.Black, false, b.EnPassantFile) {
	//	graphics.PrintBinaryBoard(move)
	//}
	//
	//fmt.Println(b.Evaluate())
	//
	//undo := b.ApplyMove(0)
	//graphics.PrintBinaryBoard(b.White)
	//undo()
	//graphics.PrintBinaryBoard(b.White)

}
