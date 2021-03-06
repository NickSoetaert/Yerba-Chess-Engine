package moveGen

import (
	"fmt"
	"math"
	"math/bits"
)

var BlackCount = 0 //used for debug
var WhiteCount = 0 //used for debug

func (b *Board) MiniMax(plyLeft int, alpha, beta float64) float64 {


	//fmt.Printf("Iteration: %v\n", Iter)
	//Iter++
	//PrintBoard(*b)

	eval := b.Evaluate()
	if plyLeft == 0 {
		return eval
	}
	if math.Abs(eval) > 5000 { //TODO: make real checkmate
		fmt.Printf("checkmate - eval: %v\n", eval)
		fmt.Printf("num white kings: %v\n", bits.OnesCount64(b.WhitePieces&b.Kings))
		fmt.Printf("num black kings: %v\n", bits.OnesCount64(b.BlackPieces&b.Kings))
		fmt.Printf("black king on e1: %v\n", b.BlackPieces&b.Kings&E1 != 0)
		fmt.Printf("white king on e1: %v\n", b.WhitePieces&b.Kings&E1 != 0)
		fmt.Printf("black bishop on e1: %v\n", b.BlackPieces&b.Bishops&E1 != 0)
		fmt.Printf("ply left: %v\n", plyLeft)
		//PrintBoard(*b)
		//os.Exit(0)
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

//Counts the number of legal moves at a given ply, used for testing
func (b *Board) CountVariationsAtPly(ply, legalMoves int, printBoard bool) int {

	ply--
	for _, move := range b.GenerateLegalMoves() {
		undo := b.ApplyMove(move)
		if printBoard {
			PrintBoard(*b)
			fmt.Println("")
		}
		if ply > 0 {
			legalMoves = b.CountVariationsAtPly(ply, legalMoves, printBoard)
		} else {
			legalMoves++ //we're counting variations, so we only want to count the end states.
		}
		undo()
	}
	return legalMoves
}
