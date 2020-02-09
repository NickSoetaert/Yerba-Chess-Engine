package main

import (
	"fmt"
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

func RookAttacks(rooks, blockers uint64) uint64 {

	var possibleAttacks uint64

	count := bits.OnesCount64(rooks)
	//count must be used, else loop range will decrement with each pass.
	for i := 0; i < count; i++ {
		singleRookPosition := uint8(bits.TrailingZeros64(rooks)) //get the position of the "last" rook
		possibleAttacks |= rookBlockerMask[singleRookPosition]           //get current squares that rook can attack
		rooks ^= 1 << singleRookPosition                                 //clear that rook for the next loop iteration
	}

	return possibleAttacks
}

func singleRookAttack(square uint8, blockers uint64) uint64 {

	//Gets a binary board of pieces that intercept with possible rook path
	blockers &= rookBlockerMask[square]


	return blockers
}

func DownFill(file uint64) uint64 {
	file |= file >> 8
	file |= file >> 16
	file |= file >> 32
	return file
}


//The database of all legal moves given any square and blocker hash
var rookDB [][]uint64
var bishopDB [][]uint64

//Pre-compute all possible attacks
func Init() {
	rookDB = make([][]uint64, 64)
	bishopDB = make([][]uint64, 64)

	//Initialize the move databases to their proper size.
	//The first dimension of the DB is always 64, because there are 64 squares on a chess board.
	//The second dimension of the DB has the same number of bits as the number of possible blockers from that square.
	//This is so that you can store a move for every possible combination of blockers for a given square.
	for i:=0; i<64; i++{
		rookDB[i] = make([]uint64, 1<<uint64(bits.OnesCount64(rookBlockerMask[i])))
		bishopDB[i] = make([]uint64, 1<<uint64(bits.OnesCount64(bishopBlockerMask[i])))
	}

	for square := range rookDB {
		for blockerPermutation := range rookDB[square] {

			rookDB[square][blockerPermutation] =2
		}
	}
}

func bishopAttacks(occ uint64, square uint8) uint64 {
	occ &= bishopBlockerMask[square]	//get all squares that block the bishop
	occ *= bishopMagic[square]			//using those unique squares, calculate the hash
	occ >>= 64-uint8(bits.OnesCount64(bishopBlockerMask[square]))		//throw away the junk.
	return bishopDB[square][occ]		//Index
}

//Calculate the single bitboard of all legal rook moves given a single position.
//This function is only ever called once for each index of rookDB[][]
func slowCalcRookMoves (square int, blockers uint64) uint64 {

	var moves uint64

	//uint64 representation of origin square shifted up one rank
	up := uint64(1) << uint8(square) << 8
	for {	//calculate moves sliding upwards
	fmt.Println("up")
		if bits.TrailingZeros64(up) >= 64 {
			break	//we're off the top of the board - do not add up to moves
		}
		//If we're not off the board yet, so we can legally attack up.
		//This goes before the blockers check, because we can attack occupied squares.
		moves |= up

		//Check if our current square is occupied by a blocker. If it is, break.
		if blockers & up != 0 {
			break
		}
		up <<= 8  //move up another rank
	}

	down := uint64(1) << uint8(square) >> 8
	for {
		fmt.Println("down")
		if bits.TrailingZeros64(down) >= 64 {
			break
		}
		moves |= down
		if blockers & down != 0 {
			break
		}
		down >>= 8  //move up another rank
	}

	left := uint64(1) << uint8(square) >> 1
	for {
		fmt.Println("left")
		if (left << 1) & AFile != 0 {
			break
		}
		moves |= left
		if blockers & left != 0 {
			break
		}
		left >>= 1  //move left another rank
	}

	right := uint64(1) << uint8(square) >> 1
	for {
		fmt.Println("right")
		if (right >> 1) & HFile != 0 {
			break
		}
		moves |= right
		if blockers & right != 0 {
			break
		}
		right <<= 1  //move left another rank
	}
	fmt.Println("AAAAAAA")
	PrintBinaryBoard(moves)
	return moves
}

//Calculate the single bitboard of all legal bishop moves given a single position.
//This function is only ever called once for each index of rookDB[][]
func slowCalcBishopMoves (square int, blockers uint64) (moves uint64) {

	return
}