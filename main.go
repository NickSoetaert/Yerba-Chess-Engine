//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
)

func main() {
	//fmt.Println("Hello World")
	////40==A6

	//var move moveGen.Move
	//move = (1 << 32) - ((1 << 20)-10000)
	//fmt.Printf("%032b\n",move)
	//m := utils.IsolateBitsU32(uint32(move), 26, 31)

	//fmt.Printf("%032b\n",utils.IsolateBitsU32(uint32(move), 10, 30))
	//fmt.Printf("%032b\n",utils.SetBitsU32(uint32(move), z, 31, 0b100))
	//m := utils.IsolateBitsU32(uint32(0b11101111), 26, 31)
	//fmt.Printf("%0b\n",m)
	//
	b := moveGen.SetUpCastlingBoard()
	fmt.Printf("number of legal moves from start position: %v (should be 20)", len(b.GenerateLegalMoves()))
	for _, move := range b.GenerateLegalMoves() {
		undo := b.ApplyMove(move)
		moveGen.PrintBoard(b)
		undo()
	}

	//todo: bug - I am not updating "whitePieces" and "blackPieces" on moving to said squares.
	// This can cause a piece to change colors.

	//start := time.Now()
	//ply := 2
	//x := b.MiniMax(ply, math.Inf(-1), math.Inf(1))
	//fmt.Println("Ply: ", ply)
	//fmt.Printf("Time elapsed: %v\n", time.Since(start))
	//fmt.Printf("Eval: %v\n", x)
	//fmt.Println(moveGen.BlackCount)
	//fmt.Println(moveGen.WhiteCount)
	//fmt.Printf("en passant captures: %v\n", moveGen.EpCount)
}
