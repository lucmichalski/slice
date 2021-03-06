package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseBool(t *testing.T) {
	type args struct {
		src []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []bool{},
			},
			want: []bool{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []bool{true, true, true, false, false},
			},
			want: []bool{false, false, true, true, true},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []bool{true, true, false, false},
			},
			want: []bool{false, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseBool(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseByte(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []byte{},
			},
			want: []byte{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []byte{1, 2, 3, 4, 5},
			},
			want: []byte{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []byte{1, 2, 3, 4},
			},
			want: []byte{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseByte(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseComplex128(t *testing.T) {
	type args struct {
		src []complex128
	}
	tests := []struct {
		name string
		args args
		want []complex128
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []complex128{},
			},
			want: []complex128{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []complex128{complex(1, 1), complex(2, 2), complex(3, 3), complex(4, 4), complex(5, 5)},
			},
			want: []complex128{complex(5, 5), complex(4, 4), complex(3, 3), complex(2, 2), complex(1, 1)},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []complex128{complex(1, 1), complex(2, 2), complex(3, 3), complex(4, 4)},
			},
			want: []complex128{complex(4, 4), complex(3, 3), complex(2, 2), complex(1, 1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseComplex128(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseComplex64(t *testing.T) {
	type args struct {
		src []complex64
	}
	tests := []struct {
		name string
		args args
		want []complex64
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []complex64{},
			},
			want: []complex64{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []complex64{complex(1, 1), complex(2, 2), complex(3, 3), complex(4, 4), complex(5, 5)},
			},
			want: []complex64{complex(5, 5), complex(4, 4), complex(3, 3), complex(2, 2), complex(1, 1)},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []complex64{complex(1, 1), complex(2, 2), complex(3, 3), complex(4, 4)},
			},
			want: []complex64{complex(4, 4), complex(3, 3), complex(2, 2), complex(1, 1)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseComplex64(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseFloat32(t *testing.T) {
	type args struct {
		src []float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []float32{},
			},
			want: []float32{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []float32{1.1, 2.2, 3.3, 4.4, 5.5},
			},
			want: []float32{5.5, 4.4, 3.3, 2.2, 1.1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []float32{1.1, 2.2, 3.3, 4.4},
			},
			want: []float32{4.4, 3.3, 2.2, 1.1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseFloat32(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseFloat64(t *testing.T) {
	type args struct {
		src []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []float64{},
			},
			want: []float64{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
			},
			want: []float64{5.5, 4.4, 3.3, 2.2, 1.1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []float64{1.1, 2.2, 3.3, 4.4},
			},
			want: []float64{4.4, 3.3, 2.2, 1.1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseFloat64(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []int{},
			},
			want: []int{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []int{1, 2, 3, 4},
			},
			want: []int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt16(t *testing.T) {
	type args struct {
		src []int16
	}
	tests := []struct {
		name string
		args args
		want []int16
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []int16{},
			},
			want: []int16{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []int16{1, 2, 3, 4, 5},
			},
			want: []int16{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []int16{1, 2, 3, 4},
			},
			want: []int16{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt16(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt32(t *testing.T) {
	type args struct {
		src []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []int32{},
			},
			want: []int32{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []int32{1, 2, 3, 4, 5},
			},
			want: []int32{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []int32{1, 2, 3, 4},
			},
			want: []int32{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt32(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt64(t *testing.T) {
	type args struct {
		src []int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []int64{},
			},
			want: []int64{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []int64{1, 2, 3, 4, 5},
			},
			want: []int64{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []int64{1, 2, 3, 4},
			},
			want: []int64{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt64(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseInt8(t *testing.T) {
	type args struct {
		src []int8
	}
	tests := []struct {
		name string
		args args
		want []int8
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []int8{},
			},
			want: []int8{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []int8{1, 2, 3, 4, 5},
			},
			want: []int8{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []int8{1, 2, 3, 4},
			},
			want: []int8{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseInt8(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseRune(t *testing.T) {
	type args struct {
		src []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []rune{},
			},
			want: []rune{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []rune{1, 2, 3, 4, 5},
			},
			want: []rune{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []rune{1, 2, 3, 4},
			},
			want: []rune{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseRune(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []string{},
			},
			want: []string{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []string{"a", "b", "c", "d", "e"},
			},
			want: []string{"e", "d", "c", "b", "a"},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []string{"a", "b", "c", "d"},
			},
			want: []string{"d", "c", "b", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUint(t *testing.T) {
	type args struct {
		src []uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []uint{},
			},
			want: []uint{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []uint{1, 2, 3, 4, 5},
			},
			want: []uint{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []uint{1, 2, 3, 4},
			},
			want: []uint{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUint(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUint16(t *testing.T) {
	type args struct {
		src []uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []uint16{},
			},
			want: []uint16{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []uint16{1, 2, 3, 4, 5},
			},
			want: []uint16{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []uint16{1, 2, 3, 4},
			},
			want: []uint16{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUint16(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUint32(t *testing.T) {
	type args struct {
		src []uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []uint32{},
			},
			want: []uint32{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []uint32{1, 2, 3, 4, 5},
			},
			want: []uint32{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []uint32{1, 2, 3, 4},
			},
			want: []uint32{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUint32(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUint64(t *testing.T) {
	type args struct {
		src []uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []uint64{},
			},
			want: []uint64{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []uint64{1, 2, 3, 4, 5},
			},
			want: []uint64{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []uint64{1, 2, 3, 4},
			},
			want: []uint64{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUint64(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUint8(t *testing.T) {
	type args struct {
		src []uint8
	}
	tests := []struct {
		name string
		args args
		want []uint8
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []uint8{},
			},
			want: []uint8{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []uint8{1, 2, 3, 4, 5},
			},
			want: []uint8{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []uint8{1, 2, 3, 4},
			},
			want: []uint8{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUint8(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseUintptr(t *testing.T) {
	type args struct {
		src []uintptr
	}
	tests := []struct {
		name string
		args args
		want []uintptr
	}{
		{
			name: "nil slice",
			args: args{
				src: nil,
			},
			want: nil,
		},
		{
			name: "empty slice",
			args: args{
				src: []uintptr{},
			},
			want: []uintptr{},
		},
		{
			name: "slice odd elements",
			args: args{
				src: []uintptr{1, 2, 3, 4, 5},
			},
			want: []uintptr{5, 4, 3, 2, 1},
		},
		{
			name: "slice with even elements",
			args: args{
				src: []uintptr{1, 2, 3, 4},
			},
			want: []uintptr{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseUintptr(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseUintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleReverseInt() {
	a := []int{1, 2, 3, 4, 5}
	a = ReverseInt(a)
	fmt.Printf("%v", a)
	// Output: [5 4 3 2 1]
}
