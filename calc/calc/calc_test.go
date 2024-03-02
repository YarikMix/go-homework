package calc

import "testing"

func TestEval(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				raw: "(1+2)*3",
			},
			want: 9,
		},
		{
			name: "Test2",
			args: args{
				raw: "(1+2)-3",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Eval(tt.args.raw); got != tt.want {
				t.Errorf("Eval() = %v, want %v", got, tt.want)
			}
		})
	}
}
