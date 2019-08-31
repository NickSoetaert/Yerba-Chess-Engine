package utils

import (
	"Yerba/board"
	"fmt"
)

//65 to 90

//todo: iterate over H8-A1, concat rows together then print all at once.
func PrintBinaryBoard(b board.BinaryBoard) {
	mask := board.A8
	fmt.Println("  ----------------------------------")
	for i := 8; i >= 1; i-- {
		fmt.Printf("%d |", i)
		for j := 1; j <= 8; j++ {
			if b == mask {
				fmt.Print(" X |")
			} else {
				fmt.Print("   |")
			}
			if j != 8 {
				mask = mask << 1
			}
			//fmt.Printf("%64b\n", mask)
		}
		mask = mask >> 15
		fmt.Println("")
		fmt.Println("  ----------------------------------")
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
