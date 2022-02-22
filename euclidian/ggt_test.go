package euclidian

import "testing"

func Test_ggt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "7",
			args: args{
				a: 210,
				b: 7,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ggt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ggt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ggt2(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "7",
			args: args{
				a: 210,
				b: 7,
			},
			want: 7,
		},
		{
			name: "14",
			args: args{
				a: 14,
				b: 210,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ggt2(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ggt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_ggt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ggt(i, i*i)
	}
}

func Benchmark_ggt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ggt2(i, i*i)
	}
}
