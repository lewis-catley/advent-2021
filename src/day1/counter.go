package main

type counter struct {
	increases    int
	decreases    int
	noDifference int
}

func (c *counter) solve(in []int) {
	prev := -1
	for i, n := range in {
		if i == 0 {
			prev = n
			continue
		}
		if n > prev {
			c.increases++
		}
		if n < prev {
			c.decreases++
		}
		if n == prev {
			c.noDifference++
		}
		prev = n
	}
}
