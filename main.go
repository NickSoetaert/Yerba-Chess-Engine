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

	//PrintBinaryBoard(72340172838076926)
	a := slowCalcRookMoves(63, FourthRank|EFile|AFile)
	PrintBinaryBoard(a)


}
