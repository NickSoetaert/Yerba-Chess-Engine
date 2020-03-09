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

	for _, move := range moveGen.GetPawnMoves(b.Pawns, b.White, b.Black, true) {
		graphics.PrintBinaryBoard(move)
	}

	fmt.Println(b.Evaluate())

	undo := b.ApplyMove(0)
	graphics.PrintBinaryBoard(b.White)
	undo()
	graphics.PrintBinaryBoard(b.White)

}