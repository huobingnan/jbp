package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

type NameAndTypeConstant struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func (n *NameAndTypeConstant) Tag() int {
	return ConstantNameAndTypeInfo
}

func (n *NameAndTypeConstant) Value() interface{} {
	return []uint16{n.NameIndex, n.DescriptorIndex}
}

func (n *NameAndTypeConstant) String() string {
	return fmt.Sprintf("<CONSTANT_NameAndType_info: @%d, @%d>", n.NameIndex, n.DescriptorIndex)
}

func (n *NameAndTypeConstant) GoString() string {
	return n.String()
}

func _newNameAndTypeConstant(r *reader.ByteCodeReader) *NameAndTypeConstant {
	ret := new(NameAndTypeConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.NameIndex = classInfoIndex
	} else {
		panic("Read field ref (class info Index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.DescriptorIndex = nameAndTypeIndex
	} else {
		panic("Read field ref (name and type Index) error")
	}
	return ret
}
