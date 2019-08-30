//Nick Soetaert
package main

import (
	"Yerba/board"
	"Yerba/utils"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(board.ARank))
	//fmt.Printf("%64b\n", board.OneFile)
	utils.PrintBinaryBoard(board.B2)
}
