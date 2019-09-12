//Nick Soetaert
package main

import (
	"Yerba/board"
	//"Yerba/utils"
	"fmt"
)

func main() {
	fmt.Println("")
	//fmt.Printf("%64b\n", board.OneFile)
	//utils.PrintBinaryBoard(board.KingSide)
	myBoard := board.SetUpBoard()

	//fmt.Printf("%c\n", board.BlackPawn)

	board.PrintBoard(myBoard)

	//utils.PrintBinaryBoard(0x0001FFFAABFAD1A2)

}
