package aocfs

import (
	"io"
	"os"
)

type File interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Writer
	Stat() (os.FileInfo, error)
	Readdir(n int) ([]os.FileInfo, error)
}

type IFileSystem interface {
	Create(string) (File, error)
	MkdirAll(path string, perm os.FileMode) error
	Open(name string) (File, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
}

type FS struct{}

//go:generate mockgen -source=iface.go -destination=./mocks/iface_mock.go -package=mocks
