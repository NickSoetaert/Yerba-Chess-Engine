package moveGen

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkBoard_GenerateLegalMoves(b *testing.B) {
	board := SetUpBoardNoPawns()
	board.GenerateLegalMoves()

}

func BenchmarkBoard_MiniMax(b *testing.B) {
	board := SetUpBoardNoPawns()
	ply := 5
	eval := board.MiniMax(ply, math.Inf(-1), math.Inf(1))
	fmt.Printf("eval: %v\n", eval)
}
