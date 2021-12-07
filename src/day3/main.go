package main

import (
	"fmt"
	"strconv"

	"github.com/lewis-catley/advent-2021/common/reader"
)

const binaryLength = 12

func main() {
	r := reader.New()
	lines, err := r.ReadFile("../../input/day3.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("The answer for part 1 is:\n%s\n", part1(lines))
	fmt.Printf("\nThe answer for part 2 is:\n%s\n", part2(lines))
}

func part1(input []string) string {
	s := newSolver()
	s.count(input)
	res := s.solve()
	gamma, err := res.getGammaValue()
	if err != nil {
		panic(err)
	}
	epsilon, err := res.getEpsilonValue()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Gamma B is: %s, val is: %d.\nEpsilon B is: %s, val is: %d.\nAnswer is: %d", res.gamma, gamma, res.epsilon, epsilon, gamma*epsilon)

}

type result struct {
	gamma   string
	epsilon string
}

func (r result) getGammaValue() (int64, error) {
	return strconv.ParseInt(r.gamma, 2, 64)
}

func (r result) getEpsilonValue() (int64, error) {
	return strconv.ParseInt(r.epsilon, 2, 64)
}

type solver struct {
	indexes []*binaryCount
}

func newSolver() *solver {
	i := make([]*binaryCount, 12)
	for j := range i {
		i[j] = &binaryCount{}
	}
	return &solver{
		indexes: i,
	}
}

func (s *solver) count(lines []string) {
	for _, line := range lines {
		for i, v := range line {
			switch v {
			case '0':
				s.indexes[i].zeroCount++
			case '1':
				s.indexes[i].oneCount++
			}
		}
	}
}

func (s *solver) solve() *result {
	r := &result{}
	for _, bc := range s.indexes {
		r.gamma += bc.getGreaterValue()
		r.epsilon += bc.getLesserValue()
	}
	return r
}
