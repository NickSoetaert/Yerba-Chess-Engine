package uci

import (
	"Yerba/moveGen"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetUserMove(b moveGen.Board) moveGen.Board {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter move: ")
	move, _ := reader.ReadString('\n')

	fmt.Println("You entered ",move)
	return b.ApplyAlgebraicMove(move)
}

func parseMoveString(move string) (start, end string, isCapture bool) {
	files := make(map[string]bool)
	for i := 1; i <= 8; i++ {
		files[strconv.Itoa(i)] = true
	}

	remainingMove := move
	for _, ch := range move {
		start += string(ch)
		remainingMove = remainingMove[1:]
		fmt.Println(string(ch))
		//if we've encountered a number (aka last character of the start square), break
		if _, ok := files[string(ch)]; ok {
			break
		}
	}

	move = remainingMove
	remainingMove = ""

	for _, ch := range move {
		if string(ch) == "x" || string(ch) == "X" {
			isCapture = true
			remainingMove = move[1:]
			continue
		}
		end += string(ch)
		remainingMove = move[1:]
	}
	return
}