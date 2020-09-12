//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
)

//todo:
// -Implement checkmate
// -Optimize undo move function
// -Evaluate function
// -A BUNCH general optimization

func main() {
	fmt.Println("Hello World")

	b := moveGen.SetUpCheckmateBoard()
	moveGen.PrintBoard(b)
	fmt.Println(b.CountVariationsAtPly(6, 0, false))

	//b := moveGen.SetUpBoardNoPawns()
	//
	//for i, move := range b.GenerateLegalMoves() {
	//	undo := b.ApplyMove(move)
	//	moveGen.PrintBoard(b)
	//	undo()
	//	fmt.Println(i)
	//}

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
