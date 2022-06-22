package search

import "testing"

func Test_quickSearch(t *testing.T) {
	type args struct {
		text    string
		pattern string
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "found",
			args: args{
				text:    "1231231234",
				pattern: "1234",
			},
			want: 6,
		},
		{
			name: "found",
			args: args{
				text:    "1231231234",
				pattern: "123",
			},
			want: 0,
		},
		{
			name: "not found",
			args: args{
				text:    "1231231234",
				pattern: "12345",
			},
			want: -1,
		},
		{
			name: "not found, pattern longer than text",
			args: args{
				text:    "123",
				pattern: "1234534234132423423423",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSearch(tt.args.text, tt.args.pattern); got != tt.want {
				t.Errorf("quickSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
