package moveGen

import "math/bits"

func (b *Board) Evaluate() float64 {

	return float64(b.getMaterialImbalance())
}

func (b *Board) getMaterialImbalance() (imbalance int) {
	//white pieces
	imbalance += bits.OnesCount64(b.White & b.Pawns)
	imbalance += bits.OnesCount64(b.White&b.Knights) * 3
	imbalance += bits.OnesCount64(b.White&b.Bishops) * 3
	imbalance += bits.OnesCount64(b.White&b.Rooks) * 5
	imbalance += bits.OnesCount64(b.White&b.Queens) * 9
	//black pieces
	imbalance -= bits.OnesCount64(b.Black & b.Pawns)
	imbalance -= bits.OnesCount64(b.Black&b.Knights) * 3
	imbalance -= bits.OnesCount64(b.Black&b.Bishops) * 3
	imbalance -= bits.OnesCount64(b.Black&b.Rooks) * 5
	imbalance -= bits.OnesCount64(b.Black&b.Queens) * 9

	return
}
