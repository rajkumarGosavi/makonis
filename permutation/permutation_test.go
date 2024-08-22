package permutation

import (
	"testing"
)

func TestCheckPermutationAscii(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// fail: strings of different length
		{
			name: "fail: strings of different length",
			args: args{
				s1: "abc",
				s2: "ac",
			},
			want: false,
		},
		// fail: strings of different length
		{
			name: "fail: strings with different characters",
			args: args{
				s1: "abc",
				s2: "adc",
			},
			want: false,
		},
		// success: empty strings
		{
			name: "success: empty strings",
			args: args{
				s1: "",
				s2: "",
			},
			want: true,
		},
		// success: contains same characters
		{
			name: "success: contains same characters",
			args: args{
				s1: "abc",
				s2: "abc",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPermutationAscii(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("CheckPermutationAscii() = %v, want %v", got, tt.want)
			}
		})
	}
}
