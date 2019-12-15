//Nick Soetaert
package main

import (
	"Yerba/board"
	"Yerba/utils"
	"fmt"
)

func main() {
	fmt.Println("")
	//fmt.Printf("%64b\n", board.OneFile)

	//fmt.Printf("0x%016x\n", x)

	x := board.SetUpBoard()

	k := x.Knights & x.Black

	utils.PrintBinaryBoard(board.KnightAttacks(k))

}
