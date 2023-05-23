package minimumwindowsubstring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ADOBECODEBANC",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "a",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "aaa",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "ab",
			args: args{
				s: "ab",
				t: "a",
			},
			want: "a",
		},
		{

			name: "cabwefgewcwaefgcf",
			args: args{
				s: "cabwefgewcwaefgcf",
				t: "cae",
			},
			want: "cwae",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
