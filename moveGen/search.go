package moveGen

import (
	"math"
)

var BlackCount = 0 //used for debug
var WhiteCount = 0 //used for debug

func (b *Board) MiniMax(plyLeft int, alpha, beta float64) float64 {
	eval := b.Evaluate()
	if plyLeft == 0 {
		return eval
	}
	if math.Abs(eval) > 5000 { //TODO: make real checkmate
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
