package main

type binaryCount struct {
	zeroCount int
	oneCount  int
}

func (bc binaryCount) getGreaterValue() string {
	if bc.zeroCount > bc.oneCount {
		return "0"
	}
	return "1"
}

func (bc binaryCount) getLesserValue() string {
	if bc.zeroCount < bc.oneCount {
		return "0"
	}
	return "1"
}
