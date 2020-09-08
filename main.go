//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	//"Yerba/utils"
	"fmt"
)

func main() {
	//fmt.Println("Hello World")

	b := moveGen.SetUpBoard()

	fmt.Println(b.CountVariationsAtPly(3, 0, false))

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
