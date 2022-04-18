package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type MethodHandleConstant struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (m *MethodHandleConstant) Tag() int {
	return ConstantMethodHandleInfo
}

func (m *MethodHandleConstant) Value() interface{} {
	return []uint16{uint16(m.ReferenceKind), m.ReferenceIndex}
}

func (m *MethodHandleConstant) String() string {
	return fmt.Sprintf("<CONSTANT_MethodHandle_info: %d, @%d>", m.ReferenceKind, m.ReferenceIndex)
}

func (m *MethodHandleConstant) GoString() string {
	return m.String()
}

func _newMethodHandleConstant(r *reader.ByteCodeReader) *MethodHandleConstant {
	ret := new(MethodHandleConstant)
	if kind, ok := r.ReadU1(); ok {
		ret.ReferenceKind = kind
	} else {
		panic("Read method handle constant (reference kind) error")
	}
	if index, ok := r.ReadU2(); ok {
		ret.ReferenceIndex = index
	} else {
		panic("Read method handle constant (reference Index) error")
	}
	return ret
}
