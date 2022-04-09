package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type MethodHandleConstant struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (m *MethodHandleConstant) Tag() int {
	return ConstantMethodHandleInfo
}

func (m *MethodHandleConstant) Value() interface{} {
	return []uint16{uint16(m.referenceKind), m.referenceIndex}
}

func (m *MethodHandleConstant) String() string {
	return fmt.Sprintf("<CONSTANT_MethodHandle_info: %d, @%d>", m.referenceKind, m.referenceIndex)
}

func (m *MethodHandleConstant) GoString() string {
	return m.String()
}

func _newMethodHandleConstant(r *reader.ByteCodeReader) *MethodHandleConstant {
	ret := new(MethodHandleConstant)
	if kind, ok := r.ReadU1(); ok {
		ret.referenceKind = kind
	} else {
		panic("Read method handle constant (reference kind) error")
	}
	if index, ok := r.ReadU2(); ok {
		ret.referenceIndex = index
	} else {
		panic("Read method handle constant (reference index) error")
	}
	return ret
}
