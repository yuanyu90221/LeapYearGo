package leap_year

import "testing"

func Test_isLeapYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test with leapYear divided 400",
			args: args{year: 2000},
			want: true,
		},
		{
			name: "test with Not leapYear divided 100",
			args: args{year: 2100},
			want: false,
		},
		{
			name: "test with Not leapYear divided 4",
			args: args{year: 2004},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLeapYear(tt.args.year); got != tt.want {
				t.Errorf("isLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
