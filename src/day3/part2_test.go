package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_scrubber_getValueCount(t *testing.T) {
	commonInp := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	type fields struct {
		input []string
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   binaryCount
	}{
		{
			name: "Successful test",
			fields: fields{
				input: commonInp,
			},
			args: args{
				i: 0,
			},
			want: binaryCount{zeroCount: 5, oneCount: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scrubber{
				input: tt.fields.input,
			}
			if got := s.getValueCount(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scrubber.getValueCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scrubber_updateInput(t *testing.T) {
	oxInp := []string{
		"10110",
		"10111",
		"10101",
		"10000",
	}
	coInp := []string{
		"00100",
		"01111",
		"00111",
		"00010",
		"01010",
	}
	type fields struct {
		input      []string
		comparison func(int, int) byte
	}
	type args struct {
		i      int
		target byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "Oxygen solver",
			fields: fields{
				input:      oxInp,
				comparison: oxygenSolver,
			},
			args: args{
				i:      3,
				target: '1',
			},
			want: []string{
				"10110",
				"10111",
				"10101",
			},
		},
		{
			name: "CO2 solver",
			fields: fields{
				input:      coInp,
				comparison: coSolver,
			},
			args: args{
				i:      1,
				target: '1',
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scrubber{
				input:      tt.fields.input,
				comparison: tt.fields.comparison,
			}
			s.updateInput(tt.args.i, tt.args.target)
			assert.Equal(t, tt.want, s.input)
		})
	}
}
