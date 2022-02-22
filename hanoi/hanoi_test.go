package hanoi

import "testing"

func Test_hanoi(t *testing.T) {
	type args struct {
		from string
		via  string
		to   string
		n    int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "hanoi",
			args: args{
				from: "A",
				via:  "B",
				to:   "C",
				n:    10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hanoi(tt.args.from, tt.args.via, tt.args.to, tt.args.n)
		})
	}
}
