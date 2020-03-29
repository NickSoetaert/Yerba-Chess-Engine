package moveGen

import (
	"math"
)

var depth = 0
var bestSoFar = 0.0

func (b *Board) Search() (boards []Board) {
	var undo UndoMove
	depth++
	for _, move := range b.GenerateLegalMoves() {
		undo = b.ApplyMove(move)
		if b.Evaluate() > bestSoFar {
			bestSoFar = b.Evaluate()
		}
		boards = append(boards, *b)
		if depth < 2 {
			boards = append(boards, b.Search()...)
		}
		undo()
	}
	depth--

	return boards
}
var BlackCount = 0
var WhiteCount = 0
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