package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type MethodRefConstant struct {
	classInfoIndex, nameAndTypeInfoIndex uint16
}

func (m *MethodRefConstant) Tag() int {
	return ConstantMethodRefInfo
}

func (m *MethodRefConstant) Value() interface{} {
	return []uint16{m.classInfoIndex, m.nameAndTypeInfoIndex}
}

func (m *MethodRefConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Methodref_info: @%d, @%d>", m.classInfoIndex, m.nameAndTypeInfoIndex)
}

func _newMethodRefConstant(r *reader.ByteCodeReader) *MethodRefConstant {
	ret := new(MethodRefConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.classInfoIndex = classInfoIndex
	} else {
		panic("Read method ref (class info index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.nameAndTypeInfoIndex = nameAndTypeIndex
	} else {
		panic("Read method ref (name and type index) error")
	}
	return ret
}
