package main

import (
	"fmt"

	"github.com/lewis-catley/advent-2021/common/reader"
)

func part1() string {
	r := reader.New()
	lines, err := r.ReadFile("../../input/day1/part1.txt")
	if err != nil {
		panic(err)
	}
	input, err := convertStrToInt(lines)
	if err != nil {
		panic(err)
	}
	c := &counter{}
	c.solve(input)
	return fmt.Sprintf("Increases: %d, Decreases: %d, No Difference: %d\n", c.increases, c.decreases, c.noDifference)
}

func part2() string {
	r := reader.New()
	lines, err := r.ReadFile("../../input/day1/part1.txt")
	if err != nil {
		panic(err)
	}
	input, err := convertStrToInt(lines)
	if err != nil {
		panic(err)
	}
	slides := slidingScale(input)
	c := &counter{}
	c.solve(slides)
	return fmt.Sprintf("Increases: %d, Decreases: %d, No Difference: %d\n", c.increases, c.decreases, c.noDifference)
}
