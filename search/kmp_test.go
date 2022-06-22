package search

import (
	"reflect"
	"testing"
)

func Test_initNext(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "101011",
			args: args{
				pattern: "101011",
			},
			want: []int{
				-1, 0, 0, 1, 2, 3, 1,
			},
		},
		{
			name: "EISBEIN",
			args: args{
				pattern: "EISBEIN",
			},
			want: []int{
				-1, 0, 0, 0, 0, 1, 2, 0,
			},
		},
		{
			name: "aaabaaac",
			args: args{
				pattern: "aaabaaac",
			},
			want: []int{
				-1, 0, 1, 2, 0, 1, 2, 3, 0,
			},
		},
		{
			name: "abcxxxabcy",
			args: args{
				pattern: "abcxxxabcy",
			},
			want: []int{
				-1, 0, 0, 0, 0, 0, 0, 1, 2, 3, 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initNext(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initNext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_kmpSearch(t *testing.T) {
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
			if got := KMPSearch(tt.args.text, tt.args.pattern); got != tt.want {
				t.Errorf("KMPSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
