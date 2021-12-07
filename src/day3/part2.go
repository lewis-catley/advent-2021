package main

import (
	"fmt"
	"strconv"
)

func oxygenSolver(zeroCount, oneCount int) byte {
	if oneCount > zeroCount {
		return '1'
	}
	if oneCount == zeroCount {
		return '1'
	}
	return '0'
}

func coSolver(zeroCount, oneCount int) byte {
	if oneCount > zeroCount {
		return '0'
	}
	if oneCount == zeroCount {
		return '0'
	}
	return '1'
}

func part2(input []string) string {
	oxScrub := newScubber(input, oxygenSolver)
	oxB := oxScrub.solve()
	coScrub := newScubber(input, coSolver)
	coB := coScrub.solve()

	oxVal, err := strconv.ParseInt(oxB, 2, 64)
	if err != nil {
		panic(err)
	}
	coVal, err := strconv.ParseInt(coB, 2, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Oxygen B is: %s, val is %d.\nCO2 B is: %s, val is %d.\nAnswer is: %d", oxB, oxVal, coB, coVal, oxVal*coVal)
}

type scrubber struct {
	input      []string
	comparison func(int, int) byte
}

func newScubber(in []string, c func(int, int) byte) *scrubber {
	return &scrubber{
		input:      in,
		comparison: c,
	}
}

func (s *scrubber) solve() string {
	for i := 0; i < 12; i++ {
		bc := s.getValueCount(i)
		targetVal := s.comparison(bc.zeroCount, bc.oneCount)
		s.updateInput(i, targetVal)
		if len(s.input) == 1 {
			return s.input[0]
		}
	}
	return s.input[0]
}

func (s *scrubber) getValueCount(i int) binaryCount {
	bc := binaryCount{}
	for _, value := range s.input {
		switch v := value[i]; v {
		case '0':
			bc.zeroCount++
		case '1':
			bc.oneCount++
		}
	}
	return bc
}

func (s *scrubber) updateInput(i int, target byte) {
	out := []string{}
	for _, line := range s.input {
		if line[i] == target {
			out = append(out, line)
		}
	}
	s.input = out
}
