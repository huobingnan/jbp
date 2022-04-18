package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

type InterfaceMethodRefConstant struct {
	ClassInfoIndex       uint16
	NameAndTypeInfoIndex uint16
}

func (i *InterfaceMethodRefConstant) Tag() int {
	return ConstantInterfaceMethodRefInfo
}

func (i *InterfaceMethodRefConstant) Value() interface{} {
	return []uint16{i.ClassInfoIndex, i.NameAndTypeInfoIndex}
}

func (i *InterfaceMethodRefConstant) String() string {
	return fmt.Sprintf("<CONSTANT_InterfaceMethodref_info: @%d @%d>", i.ClassInfoIndex, i.NameAndTypeInfoIndex)
}

func (i *InterfaceMethodRefConstant) GoString() string {
	return i.String()
}

func _newInterfaceMethodRefConstant(r *reader.ByteCodeReader) *InterfaceMethodRefConstant {
	ret := new(InterfaceMethodRefConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.ClassInfoIndex = classInfoIndex
	} else {
		panic("Read interface method ref (class info Index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.NameAndTypeInfoIndex = nameAndTypeIndex
	} else {
		panic("Read interface method ref (name and type Index) error")
	}
	return ret
}
