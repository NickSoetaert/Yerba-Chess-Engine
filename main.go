//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"Yerba/utils"
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	////40==A6

	var move moveGen.Move
	move = (1 << 32) - 1
	fmt.Printf("%032b\n",move)

	x := utils.ClearBitsU32(uint32(move), 0, 31)
	fmt.Printf("%032b\n",x)
	//b := moveGen.SetUpBoard()
	//fmt.Printf("number of legal moves from start position: %v (should be 20)", len(b.GenerateLegalMoves()))
	//for i, move := range b.GenerateLegalMoves() {
	//	fmt.Printf("%032b\n",move)
	//	fmt.Println(i)
	//	b.ApplyMove(move)
	//	graphics.PrintBoard(b)
	//	//undo()
	//}

	//todo: m.getDestOccupancyAfterMove and such

	//start := time.Now()
	//ply := 6
	//x := b.MiniMax(ply, math.Inf(-1), math.Inf(1))
	//fmt.Println("Ply: ", ply)
	//fmt.Printf("Time elapsed: %v\n", time.Since(start))
	//fmt.Printf("Eval: %v\n", x)
	//fmt.Println(moveGen.BlackCount)
	//fmt.Println(moveGen.WhiteCount)
}
