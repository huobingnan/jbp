package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type NameAndTypeConstant struct {
	nameIndex, descriptorIndex uint16
}

func (n *NameAndTypeConstant) Tag() int {
	return ConstantNameAndTypeInfo
}

func (n *NameAndTypeConstant) Value() interface{} {
	return []uint16{n.nameIndex, n.descriptorIndex}
}

func (n *NameAndTypeConstant) String() string {
	return fmt.Sprintf("<CONSTANT_NameAndType_info: @%d, @%d>", n.nameIndex, n.descriptorIndex)
}

func (n *NameAndTypeConstant) GoString() string {
	return n.String()
}

func NewNameAndTypeConstant(r *reader.ByteCodeReader) *NameAndTypeConstant {
	ret := new(NameAndTypeConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.nameIndex = classInfoIndex
	} else {
		panic("Read field ref (class info index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.descriptorIndex = nameAndTypeIndex
	} else {
		panic("Read field ref (name and type index) error")
	}
	return ret
}
