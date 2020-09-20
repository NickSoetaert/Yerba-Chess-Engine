package moveGen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func BenchmarkBoard_GenerateLegalMoves(b *testing.B) {
	pChan := make(chan []Move, 1)
	nChan := make(chan []Move, 1)
	bChan := make(chan []Move, 1)
	rChan := make(chan []Move, 1)
	qbChan := make(chan []Move, 1)
	qrChan := make(chan []Move, 1)
	kChan := make(chan []Move, 1)
	castleChan := make(chan []Move, 1)
	board := SetUpBoardNoPawns()
	board.GenerateLegalMoves(pChan, nChan, bChan, rChan, qbChan, qrChan, kChan, castleChan)

}

func BenchmarkBoard_MiniMax(b *testing.B) {
	board := SetUpBoardNoPawns()
	ply := 5
	eval := board.MiniMax(ply, math.Inf(-1), math.Inf(1))
	fmt.Printf("eval: %v\n", eval)
}

func TestBoard_GenerateLegalMoves(t *testing.T) {
	pChan := make(chan []Move, 1)
	nChan := make(chan []Move, 1)
	bChan := make(chan []Move, 1)
	rChan := make(chan []Move, 1)
	qbChan := make(chan []Move, 1)
	qrChan := make(chan []Move, 1)
	kChan := make(chan []Move, 1)
	castleChan := make(chan []Move, 1)

	//Test that from starting position, we have 20 moves.
	b := SetUpBoard()
	assert.Equal(t, 20, len(b.GenerateLegalMoves(pChan, nChan, bChan, rChan, qbChan, qrChan, kChan, castleChan)))

	b = SetUpBoardNoPawns()
	assert.Equal(t, 50, len(b.GenerateLegalMoves(pChan, nChan, bChan, rChan, qbChan, qrChan, kChan, castleChan))) //would be 51 if enemy queen wasn't blocking D2

	b = SetUpWhiteCastlingBoard()
	assert.Equal(t, 26, len(b.GenerateLegalMoves(pChan, nChan, bChan, rChan, qbChan, qrChan, kChan, castleChan)))
}

//https://en.wikipedia.org/wiki/Shannon_number
func TestBoard_CountLegalMovesAtPly(t *testing.T) {
	b := SetUpBoard()
	assert.Equal(t, 20, b.CountVariationsAtPly(1, 0, false))
	assert.Equal(t, 400, b.CountVariationsAtPly(2, 0, false))
	assert.Equal(t, 8902, b.CountVariationsAtPly(3, 0, false))
	assert.Equal(t, 197281, b.CountVariationsAtPly(4, 0, false))
	assert.Equal(t, 4865609, b.CountVariationsAtPly(5, 0, false))
	//assert.Equal(t, 119060324, b.CountVariationsAtPly(6, 0, false))
	//assert.Equal(t, 3195901860, b.CountVariationsAtPly(7, 0, false))
	//assert.Equal(t, 84998978956 , b.CountVariationsAtPly(8, 0, false))
}
