package main

import (
	"math/bits"
)

//KnightAttacks returns a BB of all squares that knights of a given color can attack, regardless of other pieces on the board.
func KnightAttacks(knights uint64) uint64 {

	var possibleAttacks uint64
	//Get number of squares that the knight can attack
	count := bits.OnesCount64(knights)
	//count must be used, else loop range will decrement with each pass.
	for i := 0; i < count; i++ {
		singleKnightPosition := uint8(bits.TrailingZeros64(knights))
		possibleAttacks |= knightAttacks[singleKnightPosition] //get current square that knight can attack
		knights ^= 1 << singleKnightPosition                   //now clear that knight for the next loop iteration
	}
	return possibleAttacks
}

func OldRookAttacks(rooks, blockers uint64) uint64 {

	var possibleAttacks uint64

	count := bits.OnesCount64(rooks)
	//count must be used, else loop range will decrement with each pass.
	for i := 0; i < count; i++ {
		singleRookPosition := uint8(bits.TrailingZeros64(rooks)) //get the position of the "last" rook
		possibleAttacks |= rookBlockerMask[singleRookPosition]   //get current squares that rook can attack
		rooks ^= 1 << singleRookPosition                         //clear that rook for the next loop iteration
	}

	return possibleAttacks
}




//Pre-compute all possible attacks
//TODO: use pointers
func Init() ([][]uint64, [][]uint64){
	//TODO: can I make these [][]uint16 because no mask has a popcount past 12?
	rookDB := make([][]uint64, 64)
	bishopDB := make([][]uint64, 64)

	//Initialize the move databases to their proper size.
	//The first dimension of the DB is always 64, because there are 64 squares on a chess board.
	//The second dimension of the DB has the same number of bits as the number of possible blockers from that square.
	//This is so that you can store a move for every possible combination of blockers for a given square.
	for i := 0; i < 64; i++ {
		rookDB[i] = make([]uint64, 1<<uint64(bits.OnesCount64(rookBlockerMask[i])))
		bishopDB[i] = make([]uint64, 1<<uint64(bits.OnesCount64(bishopBlockerMask[i])))
	}

	//Populate rook DB
	for square, blocker := range rookBlockerMask {	//for all 64 squares
		getPermutations(uint8(square), U8PopCount(blocker), blocker, true, rookDB)
	}

	//Populate bishop DB
	for square, blocker := range bishopBlockerMask {
		getPermutations(uint8(square), U8PopCount(blocker), blocker,false, bishopDB)
	}
	return rookDB, bishopDB
}

//Takes a mask of 1s and 0s, representing possible board occupations.
//
var count = 0
var idxs = make(map[uint64]int)
func getPermutations(square, blockersLeft uint8, mask uint64, isRook bool, db [][]uint64) {
	//if we've run through the entire mask
	if blockersLeft == 0 {
		if isRook {
			index := (mask * rookMagic[square]) >> (64 - UPopCount(rookBlockerMask[square])) //throw away the junk
			idxs[mask]++
			count++
			db[square][index] = slowCalcRookMoves(square, mask)

		} else {
			index := (mask * bishopMagic[square]) >>  (64 - UPopCount(bishopBlockerMask[square])) //throw away the junk
			db[square][index] = slowCalcBishopMoves(square, mask)

		}
		return
	}

	blockersLeft--

	getPermutations(square, blockersLeft, mask, isRook, db)

	//Calculate a mask where 1 is set to 0 at our current bit
	var currentBit uint64
	if isRook {
		currentBit = rookBlockerMask[square]
	} else {
		currentBit = bishopBlockerMask[square]
	}
	//clear ones that we've seen
	for i := 0; i < int(blockersLeft); i++ {
		currentBit &= currentBit-1
	}
	mask ^= 1<<bits.TrailingZeros64(currentBit) //clears the least significant bit

	getPermutations(square, blockersLeft, mask, isRook, db)

}


func GetRookAttacks(db [][]uint64, square uint8, allPieces uint64) uint64 {
	blockers := rookBlockerMask[square] & allPieces
	hash := blockers * rookMagic[square]
	hash >>= 64 - uint8(UPopCount(rookBlockerMask[square]))
	return db[square][hash]
}


//Calculate the single bitboard of all legal rook moves given a single position.
//This function is only ever called once for each index of rookDB[][]
func slowCalcRookMoves(square uint8, blockers uint64) uint64 {
	var moves uint64

	//uint64 representation of origin square shifted up one rank
	up := uint64(1) << square << 8
	for up != 0{
		if up & FirstRank != 0 {
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
		if down & EighthRank != 0 {
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
		if left & HFile != 0 {
			break
		}
		moves |= left
		if blockers&left != 0 {
			break
		}
		left >>= 1 //move left another rank
	}

	right := uint64(1) << square << 1
	for right != 0{
		if right & AFile != 0 {
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
		if upRight & (AFile|FirstRank) != 0{
			break
		}
		moves |= upRight
		if blockers & upRight != 0 {
			break
		}
		upRight <<= 9 //up one rank and to the right one file
	}

	upLeft := uint64(1) << square << 7
	for upLeft != 0 {
		if upLeft & (HFile|FirstRank) != 0 {
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
		if downRight & (AFile | EighthRank)!=0 {
			break
		}
		moves |= downRight
		if blockers&downRight != 0 {
			break
		}
		downRight >>= 7
	}

	downLeft := uint64(1) << square >> 9
	for downLeft !=0 {
		if downLeft & (HFile | EighthRank)!=0{
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
