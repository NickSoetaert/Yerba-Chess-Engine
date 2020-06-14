package moveGen

import (
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
	_ = board.MiniMax(ply, math.Inf(-1), math.Inf(1))
}