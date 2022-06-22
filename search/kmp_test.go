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
