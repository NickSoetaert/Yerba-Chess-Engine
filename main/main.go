//Nick Soetaert
package main

import (
	"Yerba/board"
	"Yerba/utils"
	"fmt"
)

func main() {
	fmt.Printf("%64b\n", board.A8)
	//fmt.Printf("%64b\n", board.OneFile)
	utils.PrintBinaryBoard(board.H8)
}
