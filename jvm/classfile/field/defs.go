package field

import (
	"bytecodeparser/jvm/classfile/attribute"
	"bytecodeparser/jvm/classfile/reader"
)

type Field struct {
	accessFlags     uint16                // 访问标志符
	nameIndex       uint16                // 名称在常量池中索引
	descriptorIndex uint16                // 描述符在常量池中索引
	attributes      []attribute.Attribute // 属性表
}

// New 新建一个属性
func New(r *reader.ByteCodeReader) *Field {
	ret := new(Field)
	if flags, ok := r.ReadU2(); ok {
		ret.accessFlags = flags
	} else {
		panic("Read field access flags error")
	}
	if idx, ok := r.ReadU2(); ok {
		ret.nameIndex = idx
	} else {
		panic("Read field name index error")
	}
	if desc, ok := r.ReadU2(); ok {
		ret.descriptorIndex = desc
	} else {
		panic("Read field descriptor index error")
	}
	// 读取字段的属性
	if count, ok := r.ReadU2(); ok {
		if count > 0 {
			ret.attributes = make([]attribute.Attribute, count)
		}
	} else {
		panic("Read field attribute count error")
	}
	return ret
}
