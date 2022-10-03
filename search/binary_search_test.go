package search

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		list []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "found",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				key:  3,
			},
			want: 2,
		},
		{
			name: "found 2",
			args: args{
				list: []int{-100, -2, 0, 1, 2, 3, 4, 5, 6},
				key:  0,
			},
			want: 2,
		},
		{
			name: "found 3",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6, 8, 100, 10000},
				key:  100,
			},
			want: 7,
		},
		{
			name: "not found",
			args: args{
				list: []int{1, 2, 3, 4, 5, 6},
				key:  7,
			},
			want: -1,
		},
		{
			name: "empty",
			args: args{
				list: []int{},
				key:  7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.list, tt.args.key); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
