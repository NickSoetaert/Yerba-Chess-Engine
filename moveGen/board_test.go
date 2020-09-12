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
	assert.Equal(t, 50, len(b.GenerateLegalMoves())) //would be 51 if enemy queen wasn't blocking D2

	b = SetUpCastlingTestBoard()
	assert.Equal(t, 26, len(b.GenerateLegalMoves()))
}

//https://en.wikipedia.org/wiki/Shannon_number
func TestBoard_CountLegalMovesAtPly(t *testing.T) {
	b := SetUpBoard()
	assert.Equal(t, 20, b.CountVariationsAtPly(1, 0, false))
	assert.Equal(t, 400, b.CountVariationsAtPly(2, 0, false))
	assert.Equal(t, 8902, b.CountVariationsAtPly(3, 0, false))
	assert.Equal(t, 197281, b.CountVariationsAtPly(4, 0, false))
	//assert.Equal(t, 4865609, b.CountVariationsAtPly(5, 0, false))
	//assert.Equal(t, 119060324, b.CountVariationsAtPly(6, 0, false))
	//assert.Equal(t, 3195901860, b.CountVariationsAtPly(7, 0, false))
	//assert.Equal(t, 84998978956 , b.CountVariationsAtPly(8, 0, false))
}
