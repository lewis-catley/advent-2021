package reader

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/lewis-catley/advent-2021/common/aocfs"
	"github.com/lewis-catley/advent-2021/common/aocfs/mocks"
	"github.com/stretchr/testify/assert"
)

func TestReader_ReadFile(t *testing.T) {
	type openOut struct {
		f   aocfs.File
		err error
	}
	tests := []struct {
		name    string
		oOut    openOut
		wantOut []string
		wantErr bool
	}{
		{
			name: "Open file returns error",
			oOut: openOut{
				f:   mockedFile{},
				err: errors.New("Error opening file"),
			},
			wantErr: true,
		},
		{
			name: "Returns expected value",
			oOut: openOut{
				f: mockedFile{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockFs := mocks.NewMockIFileSystem(ctrl)

			mockFs.EXPECT().Open("./test/path").Return(tt.oOut.f, tt.oOut.err)

			r := Reader{
				fs: mockFs,
			}
			gotOut, err := r.ReadFile("./test/path")
			if (err != nil) != tt.wantErr {
				t.Errorf("Reader.ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, gotOut, tt.wantOut)
		})
	}
}
