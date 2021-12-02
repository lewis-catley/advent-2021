package main

import "strconv"

func convertStrToInt(s []string) ([]int, error) {
	out := []int{}
	for _, i := range s {
		n, err := strconv.Atoi(i)
		if err != nil {
			return out, err
		}
		out = append(out, n)
	}
	return out, nil
}

func slidingScale(in []int) []int {
	out := make([]int, len(in)-2)
	for i := range in {
		if i+3 > len(in) {
			break
		}
		slide := in[i : i+3]
		total := 0
		for _, n := range slide {
			total += n
		}
		out[i] = total
	}
	return out
}
