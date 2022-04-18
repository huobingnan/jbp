package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

type MethodRefConstant struct {
	ClassInfoIndex       uint16
	NameAndTypeInfoIndex uint16
}

func (m *MethodRefConstant) Tag() int {
	return ConstantMethodRefInfo
}

func (m *MethodRefConstant) Value() interface{} {
	return []uint16{m.ClassInfoIndex, m.NameAndTypeInfoIndex}
}

func (m *MethodRefConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Methodref_info: @%d, @%d>", m.ClassInfoIndex, m.NameAndTypeInfoIndex)
}

func _newMethodRefConstant(r *reader.ByteCodeReader) *MethodRefConstant {
	ret := new(MethodRefConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.ClassInfoIndex = classInfoIndex
	} else {
		panic("Read method ref (class info Index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.NameAndTypeInfoIndex = nameAndTypeIndex
	} else {
		panic("Read method ref (name and type Index) error")
	}
	return ret
}
