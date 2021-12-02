package reader

import (
	"bufio"

	"github.com/lewis-catley/advent-2021/common/aocfs"
)

type Reader struct {
	fs aocfs.IFileSystem
}

func New() Reader {
	return Reader{
		fs: aocfs.FS{},
	}
}

// ReadFile reads the contents of a file and converts it to a string slice
func (r Reader) ReadFile(path string) (out []string, err error) {
	file, err := r.fs.Open(path)
	if err != nil {
		return out, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out, nil
}
