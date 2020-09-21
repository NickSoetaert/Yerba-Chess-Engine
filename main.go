//Nick Soetaert
package main

import (
	"Yerba/moveGen"
	"fmt"
)

//todo:
// -Find missing moves that should be legal on 5+ ply (5 ply is 4,865,165 actual vs 4,865,609 expected)
// -UCI layer
// -Optimize undo move function
// -Evaluate function
// -A BUNCH general optimization
func main() {
	fmt.Println("Hello World")

	b := moveGen.BlackKingBoard()
	moveGen.PrintBoard(b)
	fmt.Println("Possible moves:")

	b.CountVariationsAtPly(1, 0, true)
}
