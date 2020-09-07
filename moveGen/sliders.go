package moveGen

import (
	"Yerba/utils"
	"math/bits"
)

//Pre-compute all possible rook and bishop attacks
func InitSlidingPieces() ([][]uint64, [][]uint64) {
	rookDB := make([][]uint64, 64)
	bishopDB := make([][]uint64, 64)

	//Initialize the move databases to their proper size.
	//The first dimension of the DB is always 64, because there are 64 squares on a chess board.
	//The second dimension of the DB has the same number of bits as the number of possible blockers from that square.
	//This is so that you can store a move for every possible combination of blockers for a given square.
	for i := 0; i < 64; i++ {
		rookDB[i] = make([]uint64, 1<<uint64(bits.OnesCount64(RookBlockerMask[i])))
		bishopDB[i] = make([]uint64, 1<<uint64(bits.OnesCount64(BishopBlockerMask[i])))
	}

	//Populate rook DB
	for square, blocker := range RookBlockerMask { //for all 64 squares
		initSlidingPiecesHelper(uint8(square), utils.U8PopCount(blocker), blocker, true, rookDB)
	}

	//Populate bishop DB
	for square, blocker := range BishopBlockerMask {
		initSlidingPiecesHelper(uint8(square), utils.U8PopCount(blocker), blocker, false, bishopDB)
	}
	return rookDB, bishopDB
}

//Given a single blocker mask, this function will populate the legal move database for the given piece at that square.
func initSlidingPiecesHelper(square, blockersLeft uint8, mask uint64, isRook bool, db [][]uint64) {
	//if we've run through the entire mask
	if blockersLeft == 0 {
		if isRook {
			//Generate 'magic' unique hash of this position to index with - e.g. rookMovesArray[square][magicHash]
			index := (mask * rookMagic[square]) >> (64 - utils.UPopCount(RookBlockerMask[square])) //throw away the junk
			db[square][index] = slowCalcRookMoves(square, mask)
		} else {
			index := (mask * bishopMagic[square]) >> (64 - utils.UPopCount(BishopBlockerMask[square])) //throw away the junk
			db[square][index] = slowCalcBishopMoves(square, mask)
		}
		return
	}
	blockersLeft--
	//calculate all possible boards in which this square is a 1
	initSlidingPiecesHelper(square, blockersLeft, mask, isRook, db)

	var currentBit uint64
	if isRook {
		currentBit = RookBlockerMask[square]
	} else {
		currentBit = BishopBlockerMask[square]
	}
	for i := 0; i < int(blockersLeft); i++ {
		currentBit &= currentBit - 1
	}
	mask ^= 1 << bits.TrailingZeros64(currentBit) //clears the least significant bit

	//calculate all possible boards in which this square is a 0
	initSlidingPiecesHelper(square, blockersLeft, mask, isRook, db)
}

func GetUnfilteredRookAttacks(db [][]uint64, square uint8, allPieces uint64) uint64 {
	blockers := RookBlockerMask[square] & allPieces
	hash := blockers * rookMagic[square]
	hash >>= 64 - uint8(utils.UPopCount(RookBlockerMask[square]))
	return db[square][hash]
}

func GetUnfilteredBishopAttacks(db [][]uint64, square uint8, allPieces uint64) uint64 {
	blockers := BishopBlockerMask[square] & allPieces
	hash := blockers * bishopMagic[square]
	hash >>= 64 - uint8(utils.UPopCount(BishopBlockerMask[square]))
	return db[square][hash]
}

//Calculate the single bitboard of all legal rook moves given a single position.
//This function is only ever called once for each index of rookDB[][]
func slowCalcRookMoves(square uint8, blockers uint64) uint64 {
	var moves uint64

	//uint64 representation of origin square shifted up one rank
	up := uint64(1) << square << 8
	for up != 0 {
		if up&FirstRank != 0 {
			break
		}
		moves |= up
		if blockers&up != 0 {
			break
		}
		up <<= 8 //move up another rank
	}

	down := uint64(1) << square >> 8
	for down != 0 {
		if down&EighthRank != 0 {
			break
		}
		moves |= down
		if blockers&down != 0 {
			break
		}
		down >>= 8
	}

	left := uint64(1) << square >> 1
	for left != 0 {
		if left&HFile != 0 {
			break
		}
		moves |= left
		if blockers&left != 0 {
			break
		}
		left >>= 1 //move left another rank
	}

	right := uint64(1) << square << 1
	for right != 0 {
		if right&AFile != 0 {
			break
		}
		moves |= right
		if blockers&right != 0 {
			break
		}
		right <<= 1
	}
	return moves
}

//Calculate the single bitboard of all legal bishop moves given a single position.
//This function is only ever called once for each index of rookDB[][]
func slowCalcBishopMoves(square uint8, blockers uint64) uint64 {
	var moves uint64

	upRight := uint64(1) << square << 9
	for upRight != 0 { //as long as we haven't wrapped around the board
		if upRight&(AFile|FirstRank) != 0 {
			break
		}
		moves |= upRight
		if blockers&upRight != 0 {
			break
		}
		upRight <<= 9 //up one rank and to the right one file
	}

	upLeft := uint64(1) << square << 7
	for upLeft != 0 {
		if upLeft&(HFile|FirstRank) != 0 {
			break
		}
		moves |= upLeft
		if blockers&upLeft != 0 {
			break
		}
		upLeft <<= 7 //up one rank and to the left one file
	}

	downRight := uint64(1) << square >> 7
	for downRight != 0 {
		if downRight&(AFile|EighthRank) != 0 {
			break
		}
		moves |= downRight
		if blockers&downRight != 0 {
			break
		}
		downRight >>= 7
	}

	downLeft := uint64(1) << square >> 9
	for downLeft != 0 {
		if downLeft&(HFile|EighthRank) != 0 {
			break
		}
		moves |= downLeft
		if blockers&downLeft != 0 {
			break
		}
		downLeft >>= 9
	}
	return moves
}

//Calculates bishop XOR rook-like moves for a given piece set (rooks, bishops, queens)
func getSliderMoves(sliders, whitePieces, blackPieces uint64, isWhiteMove, bishopMove bool, db [][]uint64, c chan []Move, piece tileOccupancy) {
	var moves []Move
	baseMove := Move(0)
	baseMove.setMoveType(normalMove)
	baseMove.setOriginOccupancy(piece)
	baseMove.setDestOccupancyAfterMove(piece)
	if isWhiteMove {
		sliders &= whitePieces
	} else {
		sliders &= blackPieces
	}
	for sliders != 0 { //While there are any sliders left to calculate moves for
		originSquareMove := baseMove
		currentSquareNum := uint8(bits.TrailingZeros64(sliders))
		originSquareMove.setOriginFromSquare(currentSquareNum)

		var possibleAttacks uint64
		if bishopMove {
			possibleAttacks = GetUnfilteredBishopAttacks(db, currentSquareNum, blackPieces|whitePieces)
		} else {
			possibleAttacks = GetUnfilteredRookAttacks(db, currentSquareNum, blackPieces|whitePieces)
		}
		if isWhiteMove {
			possibleAttacks = possibleAttacks &^ whitePieces
		} else {
			possibleAttacks = possibleAttacks &^ blackPieces
		}
		for possibleAttacks != 0 {
			move := originSquareMove

			attack := uint8(bits.TrailingZeros64(possibleAttacks))
			move.setDestFromSquare(attack)
			moves = append(moves, move)
			possibleAttacks ^= uint64(1 << attack) //clear the attack we just processed
		}
		sliders ^= uint64(1 << currentSquareNum) //clear the slider we just processed.
	}
	c <- moves
}
