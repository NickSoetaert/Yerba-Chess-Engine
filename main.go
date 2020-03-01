//Nick Soetaert
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//40==A6

	rookDB, _ := Init()
	//fmt.Println(rookDB)
	fmt.Printf("count: %v\n", count)
	fmt.Printf("Indexes: %+v\n", len(idxs))


	board := GetRookAttacks(rookDB, 40, FifthRank)
	if 1 == 1 {
		PrintBinaryBoard(board)
	}
}
