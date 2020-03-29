package moveGen

import (
	"fmt"
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

func (b *Board) MiniMax(plyLeft int, alpha, beta float64) float64 {
	if plyLeft == 0 {
		return b.Evaluate()
	}

	if b.IsWhiteMove {
		plyLeft--

		fmt.Println("White")
		maxEval := math.Inf(-1)
		for _, move := range b.GenerateLegalMoves() {
			undo := b.ApplyMove(move)
			eval := b.MiniMax(plyLeft, alpha, beta)
			undo()
			maxEval = math.Max(maxEval, eval)
			alpha = math.Max(alpha, eval)
			if beta <= alpha { //we can prune this branch
				fmt.Println("BREAKING FOR WHITE")
				break
			}
		}
		return maxEval
	} else {
		fmt.Println("Black")
		plyLeft--

		minEval := math.Inf(1)
		for _, move := range b.GenerateLegalMoves() {
			undo := b.ApplyMove(move)
			eval := b.MiniMax(plyLeft, alpha, beta)
			undo()
			minEval = math.Min(minEval, eval)
			beta = math.Min(beta, eval)
			if beta <= alpha { //we can prune this branch
				fmt.Println("BREAKING FOR BLACK")

				break
			}
		}
	}
	return b.Evaluate()
}