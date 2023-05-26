package besttimetobuyandsellstockiv

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "k2p2,4,1",
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		}, {
			name: "k2p3,2,6,5,0,3",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
