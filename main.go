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

	board := moveGen.SetUpBoard()

	for _, i := range moveGen.GetPawnMoves(board.Pawns, board.White, board.Black, true) {
		graphics.PrintBinaryBoard(i)
	}

}