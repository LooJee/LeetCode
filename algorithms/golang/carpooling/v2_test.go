package carpooling

import "testing"

func Test_carPoolingv2(t *testing.T) {
	type args struct {
		trips    [][]int
		capacity int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "[[2,1,5],[3,5,7]]",
			args: args{
				trips:    [][]int{{2, 1, 5}, {3, 5, 7}},
				capacity: 3,
			},
			want: true,
		}, {
			name: "[[9,0,1],[3,3,7]]4",
			args: args{
				trips:    [][]int{{9, 0, 1}, {3, 3, 7}},
				capacity: 4,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carPoolingv2(tt.args.trips, tt.args.capacity); got != tt.want {
				t.Errorf("carPoolingv2() = %v, want %v", got, tt.want)
			}
		})
	}
}
