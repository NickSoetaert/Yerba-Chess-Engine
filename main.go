//Nick Soetaert
package main

import (
	"Yerba/graphics"
	"Yerba/moveGen"
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	////40==A6
	b := moveGen.SetUpBoard()
	fmt.Printf("number of legal moves from start position: %v (should be 20)", len(b.GenerateLegalMoves()))
	for _, move := range b.GenerateLegalMoves() {
		fmt.Printf("%032b\n",move)
		b.ApplyMove(move)
		graphics.PrintBoard(b)
		//undo()
	}

	//todo: setDestOccupancyBeforeMove

	//start := time.Now()
	//ply := 6
	//x := b.MiniMax(ply, math.Inf(-1), math.Inf(1))
	//fmt.Println("Ply: ", ply)
	//fmt.Printf("Time elapsed: %v\n", time.Since(start))
	//fmt.Printf("Eval: %v\n", x)
	//fmt.Println(moveGen.BlackCount)
	//fmt.Println(moveGen.WhiteCount)
}
