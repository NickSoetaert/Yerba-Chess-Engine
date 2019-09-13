//Nick Soetaert
package main

import (
	//"Yerba/board"
	"Yerba/utils"
	"fmt"
)

func main() {
	fmt.Println("")
	//fmt.Printf("%64b\n", board.OneFile)
	//utils.PrintBinaryBoard(board.KingSide)
	//myBoard := board.SetUpBoard()

	//fmt.Printf("%c\n", board.BlackPawn)

	//board.PrintBoard(myBoard)

	x := utils.GetBoardKey()

	utils.PrintBinaryBoard(x)

	fmt.Printf("0x%016x\n", x)

	//x := board.BinaryBoard(utils.GetBoardKey())

	//utils.PrintBinaryBoard(board.BishopMult[63])

	//fmt.Println(x == (board.C2 | board.B3))

	//kboard := board.Board{0, x, 0, 0, 0, 0, x, 0, true}

	//board.PrintBoard(kboard)

}
