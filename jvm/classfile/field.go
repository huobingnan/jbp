package classfile

import (
	"bytecodeparser/jvm/reader"
	"fmt"
)

type JvmClassFileField struct {
	accessFlags     uint16                  // 访问标志符
	nameIndex       uint16                  // 名称在常量池中索引
	descriptorIndex uint16                  // 描述符在常量池中索引
	attributes      []JvmClassFileAttribute // 属性表
}

func (f *JvmClassFileField) String() string {
	return fmt.Sprintf("flags: %016b, name: @%d, descriptor: @%d",
		f.accessFlags, f.nameIndex, f.descriptorIndex)
}

func (f *JvmClassFileField) GoString() string {
	return f.String()
}

// NewField 新建一个属性
func NewField(r *reader.ByteCodeReader, cp ConstantPool) *JvmClassFileField {
	ret := new(JvmClassFileField)
	if flags, ok := r.ReadU2(); ok {
		ret.accessFlags = flags
	} else {
		panic(ErrorMsgFmt("Read field error", "can't read access_flags info", r.Offset()))
	}
	if idx, ok := r.ReadU2(); ok {
		ret.nameIndex = idx
	} else {
		panic(ErrorMsgFmt("Read field error", "can't read name_index info", r.Offset()))
	}
	if desc, ok := r.ReadU2(); ok {
		ret.descriptorIndex = desc
	} else {
		panic(ErrorMsgFmt("Read field error", "can't read descriptor_index info", r.Offset()))
	}
	// 读取字段的属性
	if count, ok := r.ReadU2(); ok {
		if count > 0 {
			ret.attributes = make([]JvmClassFileAttribute, 0, count)
			for i := uint16(0); i < count; i++ {
				ret.attributes = append(ret.attributes, NewAttribute(r, cp))
			}
		}
	} else {
		panic(ErrorMsgFmt("Read field error", "can't read attribute_count info", r.Offset()))
	}
	return ret
}
