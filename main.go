//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
	"math"
	"time"
)

func main() {
	//fmt.Println("Hello World")
	////40==A6
	b := moveGen.SetUpBoard()
	//fmt.Println(len(b.GenerateLegalMoves()))
	//for _, move := range b.GenerateLegalMoves() {
	//
	//	undo := b.ApplyMove(move)
	//	graphics.PrintBoard(b)
	//	undo()
	//}

	start := time.Now()
	ply := 4
	x := b.MiniMax(ply, math.Inf(-1), math.Inf(1))
	fmt.Println("Ply: ", ply)
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
	fmt.Printf("Eval: %v\n", x)
	fmt.Println(moveGen.BlackCount)
	fmt.Println(moveGen.WhiteCount)
}
