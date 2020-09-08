//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
)

//todo: Generating incorrect number of moves at 3+ ply
// -Account for moving through check when castling
// -Implement checkmate
// -Don't count capturing a king as a legal move
// -Index out of range on getNormalKingMoves() on higher plies.
// -Account for absolutely pinned pieces
// -Optimize undo move function

func main() {
	//fmt.Println("Hello World")

	b := moveGen.SetUpBoard()

	fmt.Println(b.CountVariationsAtPly(6, 0, false))

	//start := time.Now()
	//ply := 3
	//x := b.MiniMax(ply, math.Inf(-1), math.Inf(1))
	//fmt.Println("Ply: ", ply)
	//fmt.Printf("Time elapsed: %v\n", time.Since(start))
	//fmt.Printf("Eval: %v\n", x)
	//fmt.Println(moveGen.BlackCount)
	//fmt.Println(moveGen.WhiteCount)
	//fmt.Printf("en passant captures: %v\n", moveGen.EpCount)
}
