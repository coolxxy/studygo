package golang

import (
	"reflect"
	"testing"
)

func TestTrimByte(t *testing.T) {
	type args struct {
		s  []byte
		fn func(byte) bool
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"nil",
			args{
				nil,
				func(b byte) bool {
					if b == ' ' {
						return true
					}
					return false
				},
			},
			nil,
		},
		{
			"no space",
			args{
				[]byte{'a', 'b', 'c'},
				func(b byte) bool {
					if b == ' ' {
						return true
					}
					return false
				},
			},
			[]byte{'a', 'b', 'c'},
		},
		{
			"deleteSpace",
			args{
				[]byte{'a', ' ', 'b', 'c'},
				func(b byte) bool {
					if b == ' ' {
						return true
					}
					return false
				},
			},
			[]byte{'a', 'b', 'c'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimByte(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimElem(t *testing.T) {
	type args struct {
		s  []interface{}
		fn func(interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			"nil",
			args{
				nil,
				func(a interface{}) bool {
					if v, ok := a.(int); ok {
						return v < 0
					}
					return true
				},
			},
			nil,
		},
		{
			"deleteNum",
			args{
				[]interface{}{1, 2, "abc", "def", 100},
				func(a interface{}) bool {
					if v, ok := a.(int); ok {
						return v < 0
					}
					return true
				},
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimElem(tt.args.s, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimElem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"nil",
			args{s: nil},
			nil,
		},
		{
			"no space",
			args{s: []byte{'a', 'b', 'c'}},
			[]byte{'a', 'b', 'c'},
		},
		{
			"space",
			args{s: []byte{' ', 'a', 'b', ' ', 'c', ' '}},
			[]byte{'a', 'b', 'c'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSpace(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
