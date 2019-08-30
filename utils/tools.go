package utils

import (
	"Yerba/board"
	"fmt"
)

//65 to 90
func PrintBinaryBoard(b board.BinaryBoard) {
	mask := board.H8 //1000000000000000000000000000000000000000000000000000000000000000
	fmt.Println("  ----------------------------------")
	for i := 8; i >= 1; i-- {
		fmt.Printf("%d |", i)
		for j := 1; j <= 8; j++ {
			if b == mask {
				fmt.Print(" X |")
			} else {
				fmt.Print("   |")
			}
			fmt.Printf("%64b\n", mask)
			mask = mask >> 1
		}
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
