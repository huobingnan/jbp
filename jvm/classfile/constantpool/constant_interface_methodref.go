package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type InterfaceMethodRefConstant struct {
	classInfoIndex, nameAndTypeInfoIndex uint16
}

func (i *InterfaceMethodRefConstant) Tag() int {
	return ConstantInterfaceMethodRefInfo
}

func (i *InterfaceMethodRefConstant) Value() interface{} {
	return []int{int(i.classInfoIndex), int(i.nameAndTypeInfoIndex)}
}

func (i *InterfaceMethodRefConstant) String() string {
	return fmt.Sprintf("<CONSTANT_InterfaceMethodref_info: @%d @%d>", i.classInfoIndex, i.nameAndTypeInfoIndex)
}

func (i *InterfaceMethodRefConstant) GoString() string {
	return i.String()
}

func NewInterfaceMethodRefConstant(r *reader.ByteCodeReader) *InterfaceMethodRefConstant {
	ret := new(InterfaceMethodRefConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.classInfoIndex = classInfoIndex
	} else {
		panic("Read interface method ref (class info index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.nameAndTypeInfoIndex = nameAndTypeIndex
	} else {
		panic("Read interface method ref (name and type index) error")
	}
	return ret
}
