package quicksort

import (
	"reflect"
	"testing"
)

func Test_quicksort(t *testing.T) {
	type args struct {
		a     []byte
		left  int
		right int
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
				left:  0,
				right: 4,
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
				left:  0,
				right: 9,
			},
			want: []byte{
				'a', 'e', 'f', 'g', 'j', 'p', 'q', 't', 'y', 'z',
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksort(tt.args.a, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Quicksort(t *testing.T) {
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
