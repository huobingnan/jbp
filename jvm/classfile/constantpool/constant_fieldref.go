package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type FieldRefConstant struct {
	classInfoIndex, nameAndTypeInfoIndex uint16
}

func (f *FieldRefConstant) Tag() int {
	return ConstantFieldRefInfo
}

func (f *FieldRefConstant) Value() interface{} {
	return []int{int(f.classInfoIndex), int(f.nameAndTypeInfoIndex)}
}

func (f *FieldRefConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Fieldref_info: @%d, @%d>", f.classInfoIndex, f.nameAndTypeInfoIndex)
}

func (f *FieldRefConstant) GoString() string {
	return f.String()
}

func NewFieldRefConstant(r *reader.ByteCodeReader) *FieldRefConstant {
	ret := new(FieldRefConstant)
	if classInfoIndex, ok := r.ReadU2(); ok {
		ret.classInfoIndex = classInfoIndex
	} else {
		panic("Read field ref (class info index) error")
	}
	if nameAndTypeIndex, ok := r.ReadU2(); ok {
		ret.nameAndTypeInfoIndex = nameAndTypeIndex
	} else {
		panic("Read field ref (name and type index) error")
	}
	return ret
}