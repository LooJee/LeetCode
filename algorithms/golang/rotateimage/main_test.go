package rotateimage

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "matrix = [[1,2,3],[4,5,6],[7,8,9]]",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		}, {
			name: "matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]",
			args: args{
				matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			},
			want: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		}, {
			name: "matrix = [1]",
			args: args{
				matrix: [][]int{{1}},
			},
			want: [][]int{{1}},
		}, {
			name: "matrix = [[1,2], [3,4]]",
			args: args{
				matrix: [][]int{{1, 2}, {3, 4}},
			},
			want: [][]int{{3, 1}, {4, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
			if !reflect.DeepEqual(tt.args.matrix, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.matrix, tt.want)
			}
		})
	}
}
