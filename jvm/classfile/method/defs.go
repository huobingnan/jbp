package method

import (
	"bytecodeparser/jvm/classfile/attribute"
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

func ErrorMsgFmt(body, detail string, offset uint32) string {
	return fmt.Sprintf("[ERROR]:  %s (%s) @%d", body, detail, offset)
}

type Method struct {
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []attribute.JvmClassFileAttribute
}

func (m *Method) String() string {
	return fmt.Sprintf("flags: %016b, name: @%d, descriptor: @%d", m.accessFlags, m.nameIndex, m.descriptorIndex)
}

func (m *Method) GoString() string {
	return m.String()
}

func New(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *Method {
	var ok bool
	ret := new(Method)
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
	ret.attributes = make([]attribute.JvmClassFileAttribute, 0, attributeCount)
	for i := uint16(0); i < attributeCount; i++ {
		ret.attributes = append(ret.attributes, attribute.New(r, cp))
	}
	return ret
}
