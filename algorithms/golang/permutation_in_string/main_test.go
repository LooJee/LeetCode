package permutationinstring

import "testing"

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "abeidbaooo",
			args: args{
				s1: "ab",
				s2: "eidbaooo",
			},
			want: true,
		}, {
			name: "abeidboaoo",
			args: args{
				s1: "ab",
				s2: "eidboaoo",
			},
			want: false,
		}, {
			name: "helloooolleoooleh",
			args: args{
				s1: "hello",
				s2: "ooolleoooleh",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
