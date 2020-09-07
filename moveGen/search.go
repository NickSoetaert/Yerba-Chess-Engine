package moveGen

import (
	"fmt"
	"math"
	"math/bits"
	"os"
)

var BlackCount = 0 //used for debug
var WhiteCount = 0 //used for debug
var Iter = 0

func (b *Board) MiniMax(plyLeft int, alpha, beta float64) float64 {
	//fmt.Printf("Iteration: %v\n", Iter)
	//Iter++
	//PrintBoard(*b)
	eval := b.Evaluate()
	if plyLeft == 0 {
		return eval
	}
	if math.Abs(eval) > 5000 { //TODO: make real checkmate
		fmt.Printf("exiting - checkmate - eval: %v\n", eval)
		fmt.Printf("num white kings: %v\n", bits.OnesCount64(b.WhitePieces&b.Kings))
		fmt.Printf("num black kings: %v\n", bits.OnesCount64(b.BlackPieces&b.Kings))
		fmt.Printf("black king on e1: %v\n",b.BlackPieces & b.Kings & E1 != 0)
		fmt.Printf("white king on e1: %v\n",b.WhitePieces & b.Kings & E1 != 0)

		PrintBoard(*b)
		os.Exit(0)
		return eval
	}

	if b.IsWhiteMove {
		WhiteCount++
		plyLeft--

		maxEval := math.Inf(-1)
		for _, move := range b.GenerateLegalMoves() {
			undo := b.ApplyMove(move)
			eval := b.MiniMax(plyLeft, alpha, beta)
			undo()
			maxEval = math.Max(maxEval, eval)
			alpha = math.Max(alpha, eval)
			if beta <= alpha { //we can prune this branch
				break
			}
		}
		return maxEval
	} else {
		BlackCount++
		plyLeft--

		minEval := math.Inf(1)
		for _, move := range b.GenerateLegalMoves() {
			undo := b.ApplyMove(move)
			eval := b.MiniMax(plyLeft, alpha, beta)
			undo()
			minEval = math.Min(minEval, eval)
			beta = math.Min(beta, eval)
			if beta <= alpha { //we can prune this branch
				break
			}
		}
	}
	return eval
}
