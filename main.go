//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")
	//40==A6
	b := moveGen.SetUpBoard()
	//b.IsWhiteMove = false
	//for _, move := range moveGen.GetPawnMoves(b.Pawns, b.White, b.Black, b.IsWhiteMove, b.EnPassantFile) {
	//	undo := b.ApplyMove(move)
	//	graphics.PrintBoard(b)
	//	undo()
	//}

	//
	x := b.MiniMax(4, math.Inf(-1), math.Inf(1))
	fmt.Println(x)
}
