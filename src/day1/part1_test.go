package main

import (
	"reflect"
	"testing"
)

func Test_slidingScale(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Sucessfully",
			args: args{
				in: []int{1, 2, 3, 4, 5, 6},
			},
			want: []int{6, 9, 12, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingScale(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("slidingScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
