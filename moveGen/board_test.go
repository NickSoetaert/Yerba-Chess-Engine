package moveGen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestBoard_GenerateLegalMoves(t *testing.T) {
	//Test that from starting position, we have 20 moves.
	b := SetUpBoard()
	assert.Equal(t, 20, len(b.GenerateLegalMoves()))

	b = SetUpBoardNoPawns()
	assert.Equal(t, 51, len(b.GenerateLegalMoves()))

	b = SetUpCastlingTestBoard()
	assert.Equal(t, 26, len(b.GenerateLegalMoves()))
}

//https://en.wikipedia.org/wiki/Shannon_number
func TestBoard_CountLegalMovesAtPly(t *testing.T) {
	b := SetUpBoard()
	assert.Equal(t, 20, b.CountVariationsAtPly(1, 0, false))
	assert.Equal(t, 400, b.CountVariationsAtPly(2, 0, false))
	assert.Equal(t, 8902, b.CountVariationsAtPly(3, 0, false))
}
