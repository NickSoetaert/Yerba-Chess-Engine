package utils

import (
	"fmt"
	"math/bits"
)

//GetBoardKey takes a board state for a single piece, and returns
//the long representation. For debug purposes only.
func GetBoardKey() uint64 {
	b := [8][8]string{
		{" ", " ", " ", " ", " ", " ", " ", " "}, //8
		{" ", " ", " ", " ", " ", " ", " ", " "}, //7
		{" ", " ", " ", " ", " ", " ", " ", " "}, //6
		{" ", " ", " ", " ", " ", " ", " ", " "}, //5
		{" ", " ", " ", " ", " ", " ", " ", " "}, //4
		{"x", "x", "x", " ", " ", " ", " ", " "}, //3
		{"x", " ", "x", " ", " ", " ", " ", " "}, //2
		{"x", "x", "x", " ", " ", " ", " ", " "}, //1
		//A    B    C    D    E    F    G    H
	}

	var result uint64
	for i := uint8(0); i < 64; i++ {
		if b[7-(i/8)][i%8] != " " {
			result += 1 << i
		}
	}
	return result << 1
}

func UPopCount(b uint64) uint {
	return uint(bits.OnesCount64(b))
}

func U8PopCount(b uint64) uint8 {
	return uint8(bits.OnesCount64(b))
}

func IsolateLsb(b uint64) uint64 {
	return 1 << bits.TrailingZeros64(b)
}

//Expects 0 <= startIndex <= endIndex <= 31, with 0 representing the most significant bit.
//will return the resulting bits in said range, not keeping their "significance".
//setting the start and end at the same index will clear the single bit at that index.
//Example(using 8 bits instead of 32):
//IsolateBits(0b11101111, 2, 7) -> 111011. Note that this is NOT 11101100 (which is keeping significance)
func IsolateBitsU32(bits, startIndex, endIndex uint32) uint32 {
	//Step one: Set all the bits to the left of startIndex to 0
	bitsToTheLeft := uint32(1 << (startIndex) - 1) //get all the bits to the left of startIndex
	bitsToTheLeft = bitsToTheLeft << (32 - startIndex) //this will align bitsToTheLeft with the original move
	bits = bits &^ bitsToTheLeft // Clear all bits to the left of startIndex
	return bits >> (31 - endIndex) // Clear bits to the right of endIndex
}

//Expects 0 <= startIndex <= endIndex <= 31, with 0 representing the most significant bit.
//First, sets all bits in the given range to 0.
//Then, starting at the least significant bit WITHIN the given range, set the cleared bits to the new bits.
//Examples(using 8 bits instead of 32):
//SetBits(0b11111111, 0, 0, 0) -> 01111111
//SetBits(0b01010101, 2, 5, 0b100) ->(clear) 01000001 -> (set) 01010001 (notice this is not 01 1000 01 but rather 01 0100 01)
//SetBits(0b11111111, 2, 7, 0b010101) -> 11010101

func SetBitsU32(oldBits, startIndex, endIndex, newBits uint32) uint32 {
	//First, clear the bits in range [startIndex, endIndex]
	oldBits = ClearBitsU32(oldBits, startIndex, endIndex)
	//Now, line up the new bits with the cleared bits
	//We need to ensure that 0s will be left filled if end-start > num bits. (What if you passed 0b001 to set 3 bits?)
	newBits = newBits << (31-endIndex)

	return oldBits | newBits
}

//Expects 0 <= startIndex <= endIndex <= 31, with 0 representing the most significant bit.
//will set the bits in said range to 0
//setting the start and end at the same index will clear the single bit at that index.
func ClearBitsU32(oldBits, startIndex, endIndex uint32) uint32 {
	//Get a block of bits n long to clear with
	block := 1 << (endIndex+1 - startIndex)-1

	//Shift these bits until they line up with the bits you wish to clear
	clear := uint32(block << (31-endIndex))

	//use those bits to clear oldBits
	return oldBits &^ clear

}

//PrintBinaryBoard takes a bitboard and prints it in chess-board format
func PrintBinaryBoard(b uint64) {
	mask := uint64(72057594037927936) //A8 or top left corner
	fmt.Println("  ---------------------------------")
	for i := 8; i >= 1; i-- {
		fmt.Printf("%d |", i)
		for j := 1; j <= 8; j++ {
			if bits.OnesCount64(b&mask) == 1 {
				fmt.Print(" X |")
			} else {
				fmt.Print("   |")
			}
			if j != 8 {
				mask = mask << 1
			}
		}
		mask = mask >> 15
		fmt.Println("")
		fmt.Println("  ---------------------------------")
	}
	//print numbers at bottom with 3 spaces of padding
	for i := 'A' - 3; i <= 'H'; i++ {
		if i >= 'A' {
			fmt.Printf(" %c |", i)
		} else if i == 'A'-1 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println("")
}
