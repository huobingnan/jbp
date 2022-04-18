package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type JvmClassFileMethod struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []JvmClassFileAttribute
}

func (m *JvmClassFileMethod) String() string {
	return fmt.Sprintf("flags: %016b, name: @%d, descriptor: @%d", m.accessFlags, m.nameIndex, m.descriptorIndex)
}

func (m *JvmClassFileMethod) GoString() string {
	return m.String()
}

func NewMethod(r *reader.ByteCodeReader, cp ConstantPool) *JvmClassFileMethod {
	var ok bool
	ret := new(JvmClassFileMethod)
	ret.accessFlags, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read method error", "can't read access_flags info", r.Offset()))
	}
	ret.nameIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read method error", "can't read name_index info", r.Offset()))
	}
	ret.descriptorIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read method error", "can't read descriptor_index info", r.Offset()))
	}
	attributeCount, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read method error", "can't read attribute_count info", r.Offset()))
	}
	ret.attributes = make([]JvmClassFileAttribute, 0, attributeCount)
	for i := uint16(0); i < attributeCount; i++ {
		ret.attributes = append(ret.attributes, NewAttribute(r, cp))
	}
	return ret
}
