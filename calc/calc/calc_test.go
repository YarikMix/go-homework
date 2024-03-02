package calc

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test sum",
			args: args{
				raw: "10+15",
			},
			want: 25,
		},
		{
			name: "Test diff",
			args: args{
				raw: "10-8",
			},
			want: 2,
		},
		{
			name: "Test multiple",
			args: args{
				raw: "10*15",
			},
			want: 150,
		},
		{
			name: "Test div",
			args: args{
				raw: "3/2",
			},
			want: 1.5,
		},
		{
			name: "Test parenthesis",
			args: args{
				raw: "(1+2)*3",
			},
			want: 9,
		},
		{
			name: "Test1",
			args: args{
				raw: "(1+2)/0",
			},
			want: math.Inf(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var value, err = Eval(tt.args.raw)
			assert.NoError(t, err)
			assert.InDelta(t, value, tt.want, 0.001)
		})
	}
}

func TestEvalFail(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test1",
			args: args{
				raw: "1++4",
			},
		},
		{
			name: "Test2",
			args: args{
				raw: "1*/4",
			},
		},
		{
			name: "Test3",
			args: args{
				raw: "(1+2))-3",
			},
		},
		{
			name: "Test4",
			args: args{
				raw: "()+10",
			},
		},
		{
			name: "Test5",
			args: args{
				raw: "15/-25",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var _, err = Eval(tt.args.raw)
			assert.Error(t, err)
		})
	}
}
