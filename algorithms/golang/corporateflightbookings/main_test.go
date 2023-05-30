package corporateflightbookings

import (
	"reflect"
	"testing"
)

func Test_corpFlightBookings(t *testing.T) {
	type args struct {
		bookings [][]int
		n        int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5",
			args: args{
				bookings: [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}},
				n:        5,
			},
			want: []int{10, 55, 45, 25, 25},
		}, {
			name: "bookings = [[1,2,10],[2,2,15]], n = 2",
			args: args{
				bookings: [][]int{{1, 2, 10}, {2, 2, 15}},
				n:        2,
			},
			want: []int{10, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := corpFlightBookings(tt.args.bookings, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("corpFlightBookings() = %v, want %v", got, tt.want)
			}
		})
	}
}
