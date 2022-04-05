package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

// DynamicConstant 一个动态计算的常量
type DynamicConstant struct {
	bootstrapMethodAttrIndex, nameAndTypeIndex uint16
}

func (d *DynamicConstant) Tag() int {
	return ConstantDynamicInfo
}

func (d *DynamicConstant) Value() interface{} {
	return []int{int(d.bootstrapMethodAttrIndex), int(d.nameAndTypeIndex)}
}

func (d *DynamicConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Dynamic_info: @%d, @%d>", d.bootstrapMethodAttrIndex, d.nameAndTypeIndex)
}

func (d *DynamicConstant) GoString() string {
	return d.String()
}

func NewDynamicConstant(r *reader.ByteCodeReader) *DynamicConstant {
	ret := new(DynamicConstant)
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
