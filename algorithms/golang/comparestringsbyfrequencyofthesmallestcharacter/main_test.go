package comparestringsbyfrequencyofthesmallestcharacter

import (
	"reflect"
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	type args struct {
		queries []string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "queries = [\"cbd\"], words = [\"zaaaz\"]",
			args: args{
				queries: []string{"cbd"},
				words:   []string{"zaaaz"},
			},
			want: []int{1},
		}, {
			name: "queries = [\"bbb\",\"cc\"], words = [\"a\",\"aa\",\"aaa\",\"aaaa\"]",
			args: args{
				queries: []string{"bbb", "cc"},
				words:   []string{"a", "aa", "aaa", "aaaa"},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSmallerByFrequency(tt.args.queries, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSmallerByFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
