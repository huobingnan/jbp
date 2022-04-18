package classfile

import (
	"bytecodeparser/jvm/reader"
	"fmt"
)

type MethodTypeConstant struct {
	DescriptorIndex uint16
}

func (m *MethodTypeConstant) Tag() int {
	return ConstantMethodTypeInfo
}

func (m *MethodTypeConstant) Value() interface{} {
	return m.DescriptorIndex
}

func (m *MethodTypeConstant) String() string {
	return fmt.Sprintf("<CONSTANT_MethodType_info: @%d>", m.DescriptorIndex)
}

func (m *MethodTypeConstant) GoString() string {
	return m.String()
}

func _newMethodTypeConstant(r *reader.ByteCodeReader) *MethodTypeConstant {
	if index, ok := r.ReadU2(); ok {
		return &MethodTypeConstant{
			DescriptorIndex: index,
		}
	}
	panic("Read method type constant error")
}
