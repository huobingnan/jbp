package classfile

import (
	"fmt"
	"github.com/huobingnan/jbp/jvm/reader"
)

// DynamicConstant 一个动态计算的常量
type DynamicConstant struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (d *DynamicConstant) Tag() int {
	return ConstantDynamicInfo
}

func (d *DynamicConstant) Value() interface{} {
	return []uint16{d.BootstrapMethodAttrIndex, d.NameAndTypeIndex}
}

func (d *DynamicConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Dynamic_info: @%d, @%d>", d.BootstrapMethodAttrIndex, d.NameAndTypeIndex)
}

func (d *DynamicConstant) GoString() string {
	return d.String()
}

func _newDynamicConstant(r *reader.ByteCodeReader) *DynamicConstant {
	ret := new(DynamicConstant)
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
