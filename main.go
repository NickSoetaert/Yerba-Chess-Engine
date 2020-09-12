//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
)

//todo:
// -Fix EP moves
// -Optimize undo move function
// -Evaluate function
// -A BUNCH general optimization

func main() {
	fmt.Println("Hello World")

	b := moveGen.SetUpEPBoard()
	moveGen.PrintBoard(b)
	fmt.Println("possible moves:")

	b.CountVariationsAtPly(2, 0, true)


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
