package miceandcheese

import "testing"

func Test_miceAndCheese(t *testing.T) {
	type args struct {
		reward1 []int
		reward2 []int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "reward1 = [1,1,3,4], reward2 = [4,4,1,1], k = 2",
			args: args{
				reward1: []int{1, 1, 3, 4},
				reward2: []int{4, 4, 1, 1},
				k:       2,
			},
			want: 15,
		}, {
			name: "reward1 = [1,1], reward2 = [1,1], k = 2",
			args: args{
				reward1: []int{1, 1},
				reward2: []int{1, 1},
				k:       2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := miceAndCheese(tt.args.reward1, tt.args.reward2, tt.args.k); got != tt.want {
				t.Errorf("miceAndCheese() = %v, want %v", got, tt.want)
			}
		})
	}
}
