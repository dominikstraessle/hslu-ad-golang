package faculty

import "testing"

func Test_factorialRec(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "fac",
			args: args{
				n: 3628800,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialRec(tt.args.n); got != tt.want {
				t.Errorf("factorialRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_factorialIter(t *testing.T) {
	type args struct {
		n uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "fac",
			args: args{
				n: 3628800,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := factorialIter(tt.args.n); got != tt.want {
				t.Errorf("factorialRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_factorialRec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialRec(100)
	}
}
func Benchmark_factorialIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorialIter(100)
	}
}
