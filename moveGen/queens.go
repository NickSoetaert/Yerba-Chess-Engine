package moveGen

func (b *Board) unfilteredQueenAttacks(square uint8) uint64 {
	return GetUnfilteredRookAttacks(b.RookDB, square, b.White|b.Black) |
		GetUnfilteredBishopAttacks(b.BishopDB, square, b.White|b.Black)
}
