package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewis-catley/advent-2021/common/reader"
)

func main() {
	fmt.Println("Part 1 answer:", part1())
	fmt.Println("Part 2 answer:", part2())
}

func part1() string {
	r := reader.New()
	lines, err := r.ReadFile("../../input/day2.txt")
	if err != nil {
		panic(err)
	}
	commands, err := convertInput(lines)
	if err != nil {
		panic(err)
	}
	fd := calculateDestination(commands)
	return fmt.Sprintf("Depth: %d, Horizontal: %d, Total: %d", fd.depth, fd.horizontal, fd.horizontal*fd.depth)
}

func part2() string {
	r := reader.New()
	lines, err := r.ReadFile("../../input/day2.txt")
	if err != nil {
		panic(err)
	}
	commands, err := convertInput(lines)
	if err != nil {
		panic(err)
	}
	fa := calculateAim(commands)
	return fmt.Sprintf("Depth: %d, Horizontal: %d, Aim: %d, Total: %d", fa.depth, fa.horizontal, fa.aim, fa.horizontal*fa.depth)
}

type direction int

const (
	FORWARD direction = iota
	DOWN
	UP
)

type command struct {
	d direction
	u int
}

func convertInput(lines []string) ([]command, error) {
	commands := make([]command, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		c := command{}
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			return commands, nil
		}
		c.u = i

		switch d := parts[0]; d {
		case "forward":
			c.d = FORWARD
		case "down":
			c.d = DOWN
		case "up":
			c.d = UP
		}
		commands = append(commands, c)
	}
	return commands, nil
}

type finalDestination struct {
	depth      int
	horizontal int
}

func calculateDestination(cs []command) *finalDestination {
	fd := &finalDestination{}
	for _, c := range cs {
		switch d := c.d; d {
		case FORWARD:
			fd.horizontal += c.u
		case UP:
			fd.depth -= c.u
		case DOWN:
			fd.depth += c.u
		}
	}
	return fd
}

type finalAim struct {
	aim        int
	depth      int
	horizontal int
}

func calculateAim(cs []command) *finalAim {
	fa := &finalAim{}
	for _, c := range cs {
		switch d := c.d; d {
		case FORWARD:
			fa.horizontal += c.u
			fa.depth += c.u * fa.aim
		case UP:
			fa.aim -= c.u
		case DOWN:
			fa.aim += c.u
		}
	}
	return fa
}
