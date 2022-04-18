package classfile

import (
	"bytecodeparser/jvm/reader"
	"fmt"
)

type PackageConstant struct {
	NameIndex uint16
}

func (p *PackageConstant) Tag() int {
	return ConstantPackageInfo
}

func (p *PackageConstant) Value() interface{} {
	return p.NameIndex
}

func (p *PackageConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Package_info: @%d>", p.NameIndex)
}

func (p *PackageConstant) GoString() string {
	return p.String()
}

func _newPackageConstant(r *reader.ByteCodeReader) *PackageConstant {
	if index, ok := r.ReadU2(); ok {
		return &PackageConstant{index}
	}
	panic("Read package constant error")
}
