//Nick Soetaert
package main

import (
	"fmt"
)

func main() {
	fmt.Println("")
	//fmt.Printf("%64b\n", board.OneFile)

	//fmt.Printf("0x%016x\n", x)

	x := SetUpBoard()

	//k := x.Knights & x.Black

	//utils.PrintBinaryBoard(board.RookMult[0])

	PrintBoard(x)
	//z := board.DownFill(board.A4)
	//board.PrintBinaryBoard()

	//40==A6
	for i:=0; i <64; i++ {
		fmt.Println(i)
		PrintBinaryBoard(slowCalcRookMoves(i, EmptyBoard))
		fmt.Println("\n\n")
	}


}
