package maximumvalueofastringinanarray

import "testing"

func Test_maximumValue(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "strs = [\"alic3\",\"bob\",\"3\",\"4\",\"00000\"]",
			args: args{
				strs: []string{"alic3", "bob", "3", "4", "00000"},
			},
			want: 5,
		}, {
			name: "strs = [\"1\",\"01\",\"001\",\"0001\"]",
			args: args{
				strs: []string{"1", "01", "001", "0001"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumValue(tt.args.strs); got != tt.want {
				t.Errorf("maximumValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
