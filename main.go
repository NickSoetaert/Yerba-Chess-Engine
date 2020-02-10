//Nick Soetaert
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//40==A6

	rookDB, _ := Init()

	PrintBinaryBoard(GetRookAttacks(rookDB, 0, EmptyBoard))
}
