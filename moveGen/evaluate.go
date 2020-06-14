package moveGen

import "math/bits"

func (b *Board) Evaluate() float64 {

	return float64(b.getMaterialImbalance())
}

const pawnValue = 1
const knightValue = 3
const bishopValue = 3
const rookValue = 5
const queenValue = 9
const kingValue = 65536 //just a huge number that's a power of 2, because why not

func (b *Board) getMaterialImbalance() (imbalance int) {
	//white pieces
	imbalance += bits.OnesCount64(b.WhitePieces&b.Pawns) * pawnValue
	imbalance += bits.OnesCount64(b.WhitePieces&b.Knights) * knightValue
	imbalance += bits.OnesCount64(b.WhitePieces&b.Bishops) * bishopValue
	imbalance += bits.OnesCount64(b.WhitePieces&b.Rooks) * rookValue
	imbalance += bits.OnesCount64(b.WhitePieces&b.Queens) * queenValue
	imbalance += bits.OnesCount64(b.WhitePieces&b.Kings) * kingValue
	//black pieces
	imbalance -= bits.OnesCount64(b.BlackPieces&b.Pawns) * pawnValue
	imbalance -= bits.OnesCount64(b.BlackPieces&b.Knights) * knightValue
	imbalance -= bits.OnesCount64(b.BlackPieces&b.Bishops) * bishopValue
	imbalance -= bits.OnesCount64(b.BlackPieces&b.Rooks) * rookValue
	imbalance -= bits.OnesCount64(b.BlackPieces&b.Queens) * queenValue
	imbalance -= bits.OnesCount64(b.BlackPieces&b.Kings) * kingValue

	return
}
