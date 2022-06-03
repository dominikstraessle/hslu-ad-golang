package queue

import (
	"reflect"
	"testing"
)

func TestArrayRingBuffer_Offer(t *testing.T) {
	bar := &foo{'a'}
	type fields struct {
		head uint
		tail uint
		data [SIZE]*foo
	}
	type args struct {
		t *foo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "offer to empty queue",
			fields: fields{
				head: 0,
				tail: 0,
				data: [8]*foo{},
			},
			args: args{
				t: bar,
			},
			wantErr: false,
		},
		{
			name: "offer to half full queue",
			fields: fields{
				head: 3,
				tail: 1,
				data: [8]*foo{nil, bar, bar},
			},
			args: args{
				t: bar,
			},
			wantErr: false,
		},
		{
			name: "offer to full queue",
			fields: fields{
				head: 5,
				tail: 5,
				data: [8]*foo{bar, bar, bar, bar, bar, bar, bar, bar},
			},
			args: args{
				t: bar,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayRingBuffer[foo]{
				head: tt.fields.head,
				tail: tt.fields.tail,
				data: tt.fields.data,
			}
			if err := a.Offer(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Offer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestArrayRingBuffer_Poll(t *testing.T) {
	bar := &foo{'a'}

	type fields struct {
		head uint
		tail uint
		data [SIZE]*foo
	}
	tests := []struct {
		name    string
		fields  fields
		want    *foo
		wantErr bool
	}{
		{
			name: "poll from full queue",
			fields: fields{
				head: 5,
				tail: 5,
				data: [8]*foo{bar, bar, bar, bar, bar, bar, bar, bar},
			},
			want:    bar,
			wantErr: false,
		},
		{
			name: "poll from empty queue",
			fields: fields{
				head: 5,
				tail: 5,
				data: [8]*foo{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &ArrayRingBuffer[foo]{
				head: tt.fields.head,
				tail: tt.fields.tail,
				data: tt.fields.data,
			}
			got, err := a.Poll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Poll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Poll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayRingBuffer_OfferAndPoll(t *testing.T) {
	bar := &foo{'a'}

	a := &ArrayRingBuffer[foo]{
		head: 5,
		tail: 5,
		data: [8]*foo{bar, bar, bar, bar, bar, bar, bar, bar},
	}

	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if err := a.Offer(bar); err != nil {
		t.Errorf("failed to offer to %v got %v", a, err)
	}
	if got, err := a.Poll(); got != bar || err != nil {
		t.Errorf("failed to poll from %v", a)
	}
	if err := a.Offer(bar); err != nil {
		t.Errorf("failed to offer to %v got %v", a, err)
	}
}
