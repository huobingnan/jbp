package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type ModuleConstant struct {
	NameIndex uint16
}

func (m *ModuleConstant) Tag() int {
	return ConstantModuleInfo
}

func (m *ModuleConstant) Value() interface{} {
	return m.NameIndex
}

func (m *ModuleConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Module_info: @%d>", m.NameIndex)
}

func (m *ModuleConstant) GoString() string {
	return m.String()
}

func _newModuleConstant(r *reader.ByteCodeReader) *ModuleConstant {
	if index, ok := r.ReadU2(); ok {
		return &ModuleConstant{index}
	}
	panic("Read module info error")
}
