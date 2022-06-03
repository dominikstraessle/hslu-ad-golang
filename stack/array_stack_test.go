package stack

import (
	"reflect"
	"testing"
)

func TestArrayStack_Pop(t *testing.T) {
	str := new(string)
	*str = "test"
	type fields struct {
		data           [SIZE]*string
		lastEmptyIndex uint
	}
	tests := []struct {
		name    string
		fields  fields
		want    *string
		wantErr bool
	}{
		{
			name: "can pop from full stack",
			fields: fields{
				data:           [5]*string{str, str, str, str, str},
				lastEmptyIndex: 5,
			},
			want:    str,
			wantErr: false,
		},
		{
			name: "can pop from half full stack",
			fields: fields{
				data:           [5]*string{str, str, str},
				lastEmptyIndex: 3,
			},
			want:    str,
			wantErr: false,
		},
		{
			name: "cannot pop from empty stack",
			fields: fields{
				data:           [5]*string{},
				lastEmptyIndex: 0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack[string]{
				data:           tt.fields.data,
				lastEmptyIndex: tt.fields.lastEmptyIndex,
			}
			got, err := a.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_Push(t *testing.T) {
	str := new(string)
	*str = "test"
	type fields struct {
		data           [SIZE]*string
		lastEmptyIndex uint
	}
	type args struct {
		t *string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "empty stack can be pushed",
			fields: fields{
				data:           [5]*string{},
				lastEmptyIndex: 0,
			},
			args: args{
				t: str,
			},
			wantErr: false,
		},
		{
			name: "half empty stack can be pushed",
			fields: fields{
				data:           [5]*string{str, str, str, str},
				lastEmptyIndex: 4,
			},
			args: args{
				t: str,
			},
			wantErr: false,
		},
		{
			name: "full stack cannot be pushed",
			fields: fields{
				data:           [5]*string{str, str, str, str, str},
				lastEmptyIndex: 5,
			},
			args: args{
				t: str,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayStack[string]{
				data:           tt.fields.data,
				lastEmptyIndex: tt.fields.lastEmptyIndex,
			}
			if err := a.Push(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewStack(t *testing.T) {
	tests := []struct {
		name string
		want Stack[string]
	}{
		{
			name: "create new stack",
			want: &ArrayStack[string]{
				data:           [5]*string{},
				lastEmptyIndex: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStack[string](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayStack_PushAndPop(t *testing.T) {
	str := new(string)
	*str = "test"
	stack := NewStack[string]()
	if err := stack.Push(str); err != nil {
		t.Errorf("failed to push %v into %v", str, stack)
	}
	if err := stack.Push(str); err != nil {
		t.Errorf("failed to push %v into %v", str, stack)
	}
	if err := stack.Push(str); err != nil {
		t.Errorf("failed to push %v into %v", str, stack)
	}
	if err := stack.Push(str); err != nil {
		t.Errorf("failed to push %v into %v", str, stack)
	}
	if err := stack.Push(str); err != nil {
		t.Errorf("failed to push %v into %v", str, stack)
	}
	if err := stack.Push(str); err == nil {
		t.Errorf("push should fail but hasn't")
	}
	if got, err := stack.Pop(); got != str || err != nil {
		t.Errorf("failed to pop from %v, got %v and %v", stack, got, err)
	}
	if got, err := stack.Pop(); got != str || err != nil {
		t.Errorf("failed to pop from %v, got %v and %v", stack, got, err)
	}
	if got, err := stack.Pop(); got != str || err != nil {
		t.Errorf("failed to pop from %v, got %v and %v", stack, got, err)
	}
	if got, err := stack.Pop(); got != str || err != nil {
		t.Errorf("failed to pop from %v, got %v and %v", stack, got, err)
	}
	if got, err := stack.Pop(); got != str || err != nil {
		t.Errorf("failed to pop from %v, got %v and %v", stack, got, err)
	}
	if got, err := stack.Pop(); got == str || err == nil {
		t.Errorf("pop should have failed but haven't: %v, got %v and %v", stack, got, err)
	}
}
