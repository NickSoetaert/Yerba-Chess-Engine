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
	//b.IsWhiteMove = false

	fmt.Println()
	for _, move := range b.GenerateLegalMoves() {

		undo := b.ApplyMove(move)
		graphics.PrintBoard(b)
		undo()
	}

	//x := b.MiniMax(9, math.Inf(-1), math.Inf(1))
	//fmt.Println(x)
	//fmt.Println(moveGen.BlackCount)
	//fmt.Println(moveGen.WhiteCount)
}
