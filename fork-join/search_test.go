package fork_join

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "contains",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				key: 5,
			},
			want: 4,
		},
		{
			name: "does not contain",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				key: 7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchIndex(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("SearchIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchIndexParallel(t *testing.T) {
	t.Skip()
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "contains",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				key: 5,
			},
			want: 4,
		},
		{
			name: "does not contain",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				key: 7,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchIndexParallel(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("SearchIndexParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchIndexParallel2(t *testing.T) {
	t.Skip()
	var arr []int
	for i := 0; i < 100_000; i++ {
		arr = append(arr, i)
	}

	result := SearchIndexParallel(arr, 99_990)
	assert.Equal(t, 99_990, result)
}
