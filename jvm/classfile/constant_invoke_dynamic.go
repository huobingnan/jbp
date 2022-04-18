package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

type InvokeDynamicConstant struct {
	BootstrapMethodAttrIndex,
	NameAndTypeIndex uint16
}

func (i *InvokeDynamicConstant) Tag() int {
	return ConstantDynamicInfo
}

func (i *InvokeDynamicConstant) Value() interface{} {
	return []uint16{i.BootstrapMethodAttrIndex, i.NameAndTypeIndex}
}

func (i *InvokeDynamicConstant) String() string {
	return fmt.Sprintf("<CONSTANT_InvokeDynamic_info: @%d, @%d>", i.BootstrapMethodAttrIndex, i.NameAndTypeIndex)
}

func (i *InvokeDynamicConstant) GoString() string {
	return i.String()
}

func _newInvokeDynamicConstant(r *reader.ByteCodeReader) *InvokeDynamicConstant {
	ret := new(InvokeDynamicConstant)
	if bootstrapIndex, ok := r.ReadU2(); ok {
		ret.BootstrapMethodAttrIndex = bootstrapIndex
	} else {
		panic("Read dynamic constant (bootstrap method attr Index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.NameAndTypeIndex = nameAndTypeIndex
	} else {
		panic("Read dynamic constant (bootstrap method attr Index) error")
	}
	return ret
}
