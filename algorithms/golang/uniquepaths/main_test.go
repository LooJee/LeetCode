package uniquepaths

import "testing"

func Test_uniquePaths(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "m=3, n=7",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		}, {
			name: "m=3, n=2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		}, {
			name: "m=7, n=3",
			args: args{
				m: 7,
				n: 3,
			},
			want: 28,
		}, {
			name: "m=3, n=3",
			args: args{
				m: 3,
				n: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("uniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
