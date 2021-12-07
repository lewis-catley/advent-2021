package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewis-catley/advent-2021/common/reader"
)

func main() {
	r := reader.New()
	strings, err := r.ReadFile("../../input/day4.txt")
	if err != nil {
		panic(err)
	}
	bingoCalls := getInput(strings[0])
	solve(bingoCalls, strings[1:])
}

func getInput(s string) []int {
	split := strings.Split(s, ",")
	out := []int{}
	for _, intStr := range split {
		i, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}
	return out
}

func solve(calls []int, boardsInp []string) {
	boards := []*board{}
	inp := []string{}
	for _, s := range boardsInp {
		if s == "" && len(inp) != 0 {
			boards = append(boards, newBoard(inp))
		}
		if s == "" {
			inp = []string{}
			continue
		}
		inp = append(inp, s)
	}

	var winningBoard *board
	var lastWinner *board
	firstLastNo := 0
	lastLastNo := 0
	previousWinners := map[*board]interface{}{}
	for _, i := range calls {
		for _, b := range boards {
			if !b.isSolved() {
				b.valueCalled(i)
			}
			if b.isSolved() && winningBoard == nil {
				winningBoard = b
				firstLastNo = i
			}
			if _, ok := previousWinners[b]; !ok && b.isSolved() {
				previousWinners[b] = nil
				lastWinner = b
				lastLastNo = i
			}
		}
	}
	if winningBoard != nil {
		sum := winningBoard.calledSum()
		fmt.Println("This is the complete solution part 1", sum*firstLastNo)
	}
	if lastWinner != nil {
		sum := lastWinner.calledSum()
		fmt.Println("This is the complete solution part 2", sum*lastLastNo)
	}
	fmt.Println("Finished the solve")
}
