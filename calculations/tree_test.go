package calculations

import "testing"

func TestCalcCuts(t *testing.T) {
	type args struct {
		trees []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "-1 cuts v1",
			args: args{
				trees: []int{1,2,1,1,1,3,3,2},
			},
			want: -1,
		},
		{
			name: "-1 cuts v2",
			args: args{
				trees: []int{1,2,1,1,1,3,2},
			},
			want: -1,
		},
		{
			name: "0 cuts",
			args: args{
				trees: []int{1,2,1,3,2},
			},
			want: 0,
		},
		{
			name: "2 cuts v1",
			args: args{
				trees: []int{1,1,3,2},
			},
			want: 2,
		},
		{
			name: "2 cuts v2",
			args: args{
				trees: []int{3,1,2,2},
			},
			want: 2,
		},
		{
			name: "3 cuts v1",
			args: args{
				trees: []int{3,4,5,3,7},
			},
			want: 3,
		},
		{
			name: "3 cuts v2",
			args: args{
				trees: []int{7,3,5,4,3},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.trees); got != tt.want {
				t.Errorf("CalcCuts() = %v, want %v", got, tt.want)
			}
		})
	}
}

