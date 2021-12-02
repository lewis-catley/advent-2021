package reader

import "github.com/lewis-catley/advent-2021/common/aocfs"

// mockedFile useful for imitating a file in tests
type mockedFile struct {
	aocfs.File
	closeError error
}

func (f mockedFile) Close() error {
	return f.closeError
}
