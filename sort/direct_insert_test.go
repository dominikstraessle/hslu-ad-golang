package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_directInsert(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "inorder",
			args: args{
				a: []int{0, 1, 2, 3, 4, 5},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "reverse",
			args: args{
				a: []int{5, 4, 3, 2, 1, 0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name: "sort",
			args: args{
				a: []int{2, 3, 1, 4, 5, 0},
			},
			want: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := directInsert(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("directInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDirectInsert(b *testing.B) {
	var x []int
	for i := 0; i < 10000; i++ {
		x = append(x, rand.Intn(10000))
	}
	for i := 0; i < b.N; i++ {
		directInsert(x)
	}
}

func BenchmarkDirectInsertOptimistic(b *testing.B) {
	var y []int
	for i := 0; i < 10000; i++ {
		y = append(y, i)
	}
	for i := 0; i < b.N; i++ {
		directInsert(y)
	}
}

func BenchmarkDirectInsertPessimistic(b *testing.B) {
	var z []int
	for i := 0; i < 10000; i++ {
		z = append(z, 10000-i)
	}
	for i := 0; i < b.N; i++ {
		directInsert(z)
	}
}
