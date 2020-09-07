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

//expects 0 <= startIndex <= endIndex <= 31
//will return the resulting bits in said range.
func IsolateBitsU32(bits, startIndex, endIndex uint32) uint32 {
	//Step one: Set all the bits to the left of startIndex to 0
	bitsToTheLeft := uint32(1 << (startIndex) - 1) //get all the bits to the left of startIndex
	bitsToTheLeft = bitsToTheLeft << (32 - startIndex) //this will align bitsToTheLeft with the original move
	bits = bits &^ bitsToTheLeft // Clear all bits to the left of startIndex
	return bits >> (31 - endIndex) // Clear bits to the right of endIndex
}

//With the idea that newBits will be a subset of oldBits,
//sets the bits starting at startIndex with newBits
//Example(using 8 bits instead of 32):
//SetBits(01010101, 2, 6, 0100) -> 00100101
func SetBitsU32(oldBits, startIndex, endIndex, newBits uint32) uint32 {
	//First, clear the bits in range [startIndex, endIndex]
	fmt.Printf("~~~old bits: %032b\n",oldBits)
	oldBits = ClearBitsU32(oldBits, startIndex, endIndex)
	//Now, line up the new bits with the cleared bits
	//We need to ensure that 0s will be left filled if end-start > num bits. (What if you passed 0b001 to set 3 bits?)
	newBits = newBits << uint32(7-endIndex)
	fmt.Printf("~~~new bits: %032b\n", oldBits | newBits)

	return oldBits | newBits
}

//expects 0 <= startIndex <= endIndex <= 31
//will set the bits in said range to 0
func ClearBitsU32(oldBits, startIndex, endIndex uint32) uint32 {
	//Get a block of bits n long to clear with
	block := 1 << (endIndex+1 - startIndex)-1
	//Shift these bits until they line up with the bits you wish to clear
	clear := uint32(block << (7-(endIndex)))
	//use those bits to clear oldBits
	return oldBits &^ clear

}