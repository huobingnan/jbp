package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type FieldRefConstant struct {
	ClassInfoIndex       uint16
	NameAndTypeInfoIndex uint16
}

func (f *FieldRefConstant) Tag() int {
	return ConstantFieldRefInfo
}

func (f *FieldRefConstant) Value() interface{} {
	return []uint16{f.ClassInfoIndex, f.NameAndTypeInfoIndex}
}

func (f *FieldRefConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Fieldref_info: @%d, @%d>", f.ClassInfoIndex, f.NameAndTypeInfoIndex)
}

func (f *FieldRefConstant) GoString() string {
	return f.String()
}

func _newFieldRefConstant(r *reader.ByteCodeReader) *FieldRefConstant {
	ret := new(FieldRefConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.ClassInfoIndex = classInfoIndex
	} else {
		panic("Read field ref (class info Index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.NameAndTypeInfoIndex = nameAndTypeIndex
	} else {
		panic("Read field ref (name and type Index) error")
	}
	return ret
}
