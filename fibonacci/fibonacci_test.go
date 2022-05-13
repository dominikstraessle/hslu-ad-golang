package fibonacci

import "testing"

func Test_fiboRec(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "fibonacci",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "one",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				n: 30,
			},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fiboRec(tt.args.n); got != tt.want {
				t.Errorf("fiboRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibRec10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboRec(10)
	}
}

func BenchmarkFibRec30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboRec(30)
	}
}

func BenchmarkFibIter10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboIter(10)
	}
}

func BenchmarkFibIter30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboIter(30)
	}
}

func BenchmarkFibIter42(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fiboIter(42)
	}
}

func Test_fiboIter(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{

		{
			name: "fibonacci",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "one",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				n: 2,
			},
			want: 1,
		},
		{
			name: "two",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "two",
			args: args{
				n: 30,
			},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fiboIter(tt.args.n); got != tt.want {
				t.Errorf("fiboIter() = %v, want %v", got, tt.want)
			}
		})
	}
}
