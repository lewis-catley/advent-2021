package main

import (
	"strconv"
	"strings"
)

type value struct {
	isCalled bool
}

type line struct {
	values   map[int]*value
	complete bool
}

func (l *line) isComplete() bool {
	for _, v := range l.values {
		if !v.isCalled {
			return false
		}
	}
	l.complete = true
	return true
}

type board struct {
	rows     []*line
	columns  []*line
	values   map[int]*value
	complete bool
}

func newBoard(input []string) *board {
	b := &board{
		rows:    make([]*line, 5),
		columns: make([]*line, 5),
		values:  make(map[int]*value),
	}
	for y, lineIn := range input {
		splitLine := strings.Split(lineIn, " ")
		for i, s := range splitLine {
			if s == "" {
				splitLine = append(splitLine[:i], splitLine[i+1:]...)
			}
		}
		for x, valS := range splitLine {
			valInt, err := strconv.Atoi(valS)
			if err != nil {
				panic(err)
			}
			v := &value{}
			b.values[valInt] = v

			// Logic for row, and columns
			if b.rows[y] == nil {
				b.rows[y] = &line{
					values: make(map[int]*value),
				}
			}
			b.rows[y].values[valInt] = v
			if b.columns[x] == nil {
				b.columns[x] = &line{
					values: make(map[int]*value),
				}
			}
			b.columns[x].values[valInt] = v
		}
	}

	return b
}

func (b *board) valueCalled(v int) *board {
	value, ok := b.values[v]
	if !ok {
		return nil
	}
	value.isCalled = true
	return nil
}

func (b *board) isSolved() bool {
	if b.complete {
		return true
	}
	for _, r := range b.rows {
		if r.isComplete() {
			b.complete = true
			return true
		}
	}
	for _, c := range b.columns {
		if c.isComplete() {
			b.complete = true
			return true
		}
	}
	return false
}

func (b *board) calledSum() int {
	sum := 0
	for v, ok := range b.values {
		if !ok.isCalled {
			sum += v
		}
	}
	return sum
}
