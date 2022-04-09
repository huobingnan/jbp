package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type InvokeDynamicConstant struct {
	bootstrapMethodAttrIndex, nameAndTypeIndex uint16
}

func (i *InvokeDynamicConstant) Tag() int {
	return ConstantDynamicInfo
}

func (i *InvokeDynamicConstant) Value() interface{} {
	return []uint16{i.bootstrapMethodAttrIndex, i.nameAndTypeIndex}
}

func (i *InvokeDynamicConstant) String() string {
	return fmt.Sprintf("<CONSTANT_InvokeDynamic_info: @%d, @%d>", i.bootstrapMethodAttrIndex, i.nameAndTypeIndex)
}

func (i *InvokeDynamicConstant) GoString() string {
	return i.String()
}

func _newInvokeDynamicConstant(r *reader.ByteCodeReader) *InvokeDynamicConstant {
	ret := new(InvokeDynamicConstant)
	if bootstrapIndex, ok := r.ReadU2(); ok {
		ret.bootstrapMethodAttrIndex = bootstrapIndex
	} else {
		panic("Read dynamic constant (bootstrap method attr index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.nameAndTypeIndex = nameAndTypeIndex
	} else {
		panic("Read dynamic constant (bootstrap method attr index) error")
	}
	return ret
}
