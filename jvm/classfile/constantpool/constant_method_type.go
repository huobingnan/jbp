package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type MethodTypeConstant struct {
	descriptorIndex uint16
}

func (m *MethodTypeConstant) Tag() int {
	return ConstantMethodTypeInfo
}

func (m *MethodTypeConstant) Value() interface{} {
	return m.descriptorIndex
}

func (m *MethodTypeConstant) String() string {
	return fmt.Sprintf("<CONSTANT_MethodType_info: @%d>", m.descriptorIndex)
}

func (m *MethodTypeConstant) GoString() string {
	return m.String()
}

func _newMethodTypeConstant(r *reader.ByteCodeReader) *MethodTypeConstant {
	if index, ok := r.ReadU2(); ok {
		return &MethodTypeConstant{
			descriptorIndex: index,
		}
	}
	panic("Read method type constant error")
}
