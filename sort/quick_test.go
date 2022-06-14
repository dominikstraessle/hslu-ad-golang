package sort

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	x, y := getRandomizedIntSliceAndSortedSolution()
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
		{
			name: "big slice",
			args: args{
				a: x,
			},
			want: y,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Quicksort(tt.args.a))
		})
	}
}

func TestConcurrentQuicksort(t *testing.T) {
	x, y := getRandomizedIntSliceAndSortedSolution()
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
		{
			name: "big slice",
			args: args{
				a: x,
			},
			want: y,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ConcurrentQuicksort(tt.args.a))
		})
	}
}

func TestQuicksort2(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "sort",
			args: args{
				a: []byte{
					'b', 'a', 'c', 'e', 'a',
				},
			},
			want: []byte{
				'a', 'a', 'b', 'c', 'e',
			},
		},
		{
			name: "sort",
			args: args{
				a: []byte{
					'z', 'y', 'f', 'a', 'j', 'p', 'q', 'g', 'e', 't',
				},
			},
			want: []byte{
				'a', 'e', 'f', 'g', 'j', 'p', 'q', 't', 'y', 'z',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quicksort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentQuicksort2(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "sort",
			args: args{
				a: []byte{
					'b', 'a', 'c', 'e', 'a',
				},
			},
			want: []byte{
				'a', 'a', 'b', 'c', 'e',
			},
		},
		{
			name: "sort",
			args: args{
				a: []byte{
					'z', 'y', 'f', 'a', 'j', 'p', 'q', 'g', 'e', 't',
				},
			},
			want: []byte{
				'a', 'e', 'f', 'g', 'j', 'p', 'q', 't', 'y', 'z',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcurrentQuicksort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentQuicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuicksort3(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "sort",
			args: args{
				a: []byte{
					'b',
				},
			},
			want: []byte{
				'b',
			},
		},
		{
			name: "sort",
			args: args{
				a: []byte{
					'b', 'a', 'c', 'e', 'a',
				},
			},
			want: []byte{
				'a', 'a', 'b', 'c', 'e',
			},
		},
		{
			name: "sort",
			args: args{
				a: []byte{
					'z', 'y', 'f', 'a', 'j', 'p', 'q', 'g', 'e', 't',
				},
			},
			want: []byte{
				'a', 'e', 'f', 'g', 'j', 'p', 'q', 't', 'y', 'z',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Quicksort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConcurrentQuicksort3(t *testing.T) {
	type args struct {
		a []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "sort",
			args: args{
				a: []byte{
					'b',
				},
			},
			want: []byte{
				'b',
			},
		},
		{
			name: "sort",
			args: args{
				a: []byte{
					'b', 'a', 'c', 'e', 'a',
				},
			},
			want: []byte{
				'a', 'a', 'b', 'c', 'e',
			},
		},
		{
			name: "sort",
			args: args{
				a: []byte{
					'z', 'y', 'f', 'a', 'j', 'p', 'q', 'g', 'e', 't',
				},
			},
			want: []byte{
				'a', 'e', 'f', 'g', 'j', 'p', 'q', 't', 'y', 'z',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcurrentQuicksort(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcurrentQuicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuicksort500_000(b *testing.B) {
	// Aufgabe d) Laufzeiten sind immer ähnlich ca: 275335833 ns/op
	// Aufgabe f) Laufzeiten optimiert -> neu sind: 177339872 ns/op
	// Aufgabe 3.b) Laufzeiten optimiert -> neu sind: 114614168 ns/op -> quick insertion
	// 308751330
	// 127838111
	a := RandomBytes(500_000)
	for i := 0; i < b.N; i++ {
		Quicksort(a)
	}

	// byte 0-255 verschiedene Zeichen sind möglich -> bei 500000 führt dies zu sehr vielen gleichen keys
}

func BenchmarkQuicksort250_000(b *testing.B) {
	a := RandomBytes(250_000)
	for i := 0; i < b.N; i++ {
		Quicksort(a)
	}
}

func BenchmarkQuicksort125_000(b *testing.B) {
	a := RandomBytes(125_000)
	for i := 0; i < b.N; i++ {
		Quicksort(a)
	}
}

func BenchmarkQuicksort62_500(b *testing.B) {
	a := RandomBytes(62_500)
	for i := 0; i < b.N; i++ {
		Quicksort(a)
	}
}

func BenchmarkQuicksort32_250(b *testing.B) {
	a := RandomBytes(32_250)
	for i := 0; i < b.N; i++ {
		Quicksort(a)
	}
}
func BenchmarkQuicksort500(b *testing.B) {
	a := RandomBytes(500)
	for i := 0; i < b.N; i++ {
		Quicksort(a)
	}
}

func BenchmarkConcurrentQuicksort250_000(b *testing.B) {
	a := RandomBytes(250_000)
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(a)
	}
}

func BenchmarkConcurrentQuicksort125_000(b *testing.B) {
	a := RandomBytes(125_000)
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(a)
	}
}

func BenchmarkConcurrentQuicksort62_500(b *testing.B) {
	a := RandomBytes(62_500)
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(a)
	}
}

func BenchmarkConcurrentQuicksort32_250(b *testing.B) {
	a := RandomBytes(32_250)
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(a)
	}
}
func BenchmarkConcurrentQuicksort500(b *testing.B) {
	a := RandomBytes(500)
	for i := 0; i < b.N; i++ {
		ConcurrentQuicksort(a)
	}
}

func BenchmarkQuicksortStackOverflow(b *testing.B) {
	b.Skip("unable to get a stack overflow after one minute...")
	// unable to get a stack overflow after one minute...
	// java default ca 1mb
	// increase with java -Xss1048576
	a := RandomBytes(500_000_000)
	Quicksort(a)
}

// Aufgabe j) komplexität verhältnis sollte 0.000474 sein im best case
// aber average case ist ca 0.000127 was heisst, ca. 4x langsamer als best case
