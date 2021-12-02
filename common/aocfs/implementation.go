package aocfs

import "os"

func (f FS) Create(name string) (File, error) {
	return os.Create(name)
}

func (f FS) MkdirAll(path string, perm os.FileMode) error {
	return os.Mkdir(path, perm)
}

func (f FS) Open(name string) (File, error) {
	return os.Open(name)
}

func (f FS) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}
