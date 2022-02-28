package task

import "testing"

func Test_task(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "task",
			args: args{
				n: 1,
			},
		},
		{
			name: "task",
			args: args{
				n: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task(tt.args.n)
		})
	}
}

func BenchmarkTask10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		task(10)
	}
}

func BenchmarkTask100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		task(100)
	}
}

func BenchmarkTask1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		task(1000)
	}
}
