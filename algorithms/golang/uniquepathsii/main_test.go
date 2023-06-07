package uniquepathsii

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]",
			args: args{
				obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			},
			want: 2,
		}, {
			name: "obstacleGrid = [[0,1],[0,0]]",
			args: args{
				obstacleGrid: [][]int{{0, 1}, {0, 0}},
			},
			want: 1,
		}, {
			name: "obstacleGrid = [[1]]",
			args: args{
				obstacleGrid: [][]int{{1}},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
