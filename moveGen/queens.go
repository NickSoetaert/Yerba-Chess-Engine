package moveGen

func (b *Board) unfilteredQueenAttacks(square uint8) uint64 {
	return GetUnfilteredRookAttacks(b.RookDB, square, b.WhitePieces|b.BlackPieces) |
		GetUnfilteredBishopAttacks(b.BishopDB, square, b.WhitePieces|b.BlackPieces)
}
