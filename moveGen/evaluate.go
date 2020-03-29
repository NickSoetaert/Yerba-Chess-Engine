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
	imbalance += bits.OnesCount64(b.White&b.Pawns) * pawnValue
	imbalance += bits.OnesCount64(b.White&b.Knights) * knightValue
	imbalance += bits.OnesCount64(b.White&b.Bishops) * bishopValue
	imbalance += bits.OnesCount64(b.White&b.Rooks) * rookValue
	imbalance += bits.OnesCount64(b.White&b.Queens) * queenValue
	imbalance += bits.OnesCount64(b.White&b.Kings) * kingValue
	//black pieces
	imbalance -= bits.OnesCount64(b.Black&b.Pawns) * pawnValue
	imbalance -= bits.OnesCount64(b.Black&b.Knights) * knightValue
	imbalance -= bits.OnesCount64(b.Black&b.Bishops) * bishopValue
	imbalance -= bits.OnesCount64(b.Black&b.Rooks) * rookValue
	imbalance -= bits.OnesCount64(b.Black&b.Queens) * queenValue
	imbalance -= bits.OnesCount64(b.Black&b.Kings) * kingValue

	return
}
