package uniq

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	type args struct {
		lines []string
		args  Args
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test without params",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: Args{
					C: false,
					D: false,
					U: false,
					I: false,
					F: 0,
					S: 0,
				},
			},
			want: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
				"I love music of Kartik.",
			},
		},
		{
			name: "Test -c",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: Args{
					C: true,
					D: false,
					U: false,
					I: false,
					F: 0,
					S: 0,
				},
			},
			want: []string{
				"3 I love music.",
				"1 ",
				"2 I love music of Kartik.",
				"1 Thanks.",
				"2 I love music of Kartik.",
			},
		},
		{
			name: "Test -d",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: Args{
					C: false,
					D: true,
					U: false,
					I: false,
					F: 0,
					S: 0,
				},
			},
			want: []string{
				"I love music.",
				"I love music of Kartik.",
				"I love music of Kartik.",
			},
		},
		{
			name: "Test -u",
			args: args{
				lines: []string{
					"I love music.",
					"I love music.",
					"I love music.",
					"",
					"I love music of Kartik.",
					"I love music of Kartik.",
					"Thanks.",
					"I love music of Kartik.",
					"I love music of Kartik.",
				},
				args: Args{
					C: false,
					D: false,
					U: true,
					I: false,
					F: 0,
					S: 0,
				},
			},
			want: []string{
				"",
				"Thanks.",
			},
		},
		{
			name: "Test -i",
			args: args{
				lines: []string{
					"I LOVE MUSIC.",
					"I love music.",
					"I LoVe MuSiC.",
					"",
					"I love MuSIC of Kartik.",
					"I love music of kartik.",
					"Thanks.",
					"I love music of kartik.",
					"I love MuSIC of Kartik.",
				},
				args: Args{
					C: false,
					D: false,
					U: false,
					I: true,
					F: 0,
					S: 0,
				},
			},
			want: []string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Kartik.",
				"Thanks.",
				"I love music of kartik.",
			},
		},
		{
			name: "Test -f",
			args: args{
				lines: []string{
					"We love music.",
					"I love music.",
					"They love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				},
				args: Args{
					C: false,
					D: false,
					U: false,
					I: false,
					F: 1,
					S: 0,
				},
			},
			want: []string{
				"We love music.",
				"",
				"I love music of Kartik.",
				"Thanks.",
			},
		},
		{
			name: "Test -s",
			args: args{
				lines: []string{
					"I love music.",
					"A love music.",
					"C love music.",
					"",
					"I love music of Kartik.",
					"We love music of Kartik.",
					"Thanks.",
				},
				args: Args{
					C: false,
					D: false,
					U: false,
					I: false,
					F: 0,
					S: 1,
				},
			},
			want: []string{
				"I love music.",
				"",
				"I love music of Kartik.",
				"We love music of Kartik.",
				"Thanks.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solve(tt.args.lines, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
