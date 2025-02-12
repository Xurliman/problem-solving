package roman_to_int

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first",
			args: args{
				s: "MCCXXXIV",
			},
			want: 1234,
		},
		{
			name: "second",
			args: args{
				s: "DCCCXLV",
			},
			want: 845,
		},
		{
			name: "third",
			args: args{
				s: "CCXLVI",
			},
			want: 246,
		},
		{
			name: "fourth",
			args: args{
				s: "MCMXLVIII",
			},
			want: 1948,
		},
		{
			name: "fifth",
			args: args{
				s: "MCDXLIV",
			},
			want: 1444,
		},
		{
			name: "sixth",
			args: args{
				s: "MMMCMXCIX",
			},
			want: 3999,
		},
		{
			name: "seventh",
			args: args{
				s: "MMMDCCCLXXXVIII",
			},
			want: 3888,
		},
		{
			name: "eighth",
			args: args{
				s: "IX",
			},
			want: 9,
		},
		{
			name: "eighth",
			args: args{
				s: "MMCCCLXVI",
			},
			want: 2366,
		},
		{
			name: "eighth",
			args: args{
				s: "MMDCCLXXVII",
			},
			want: 2777,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
