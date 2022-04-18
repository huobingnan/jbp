package classfile

import (
	"github.com/huobingnan/jbp/jvm/reader"
)

type ConstantValueAttribute struct {
	length uint32
	Index  uint16
}

func (self *ConstantValueAttribute) Name() string { return ConstantValue }

func (self *ConstantValueAttribute) Length() uint32 { return self.length }

func _newConstantValueAttribute(r *reader.ByteCodeReader, cp ConstantPool) *ConstantValueAttribute {
	var ok bool
	ret := new(ConstantValueAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read constant value attribute error", "can't read Length info", r.Offset()))
	}
	ret.Index, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read constant value attribute error", "can't read constantvalue_index info",
			r.Offset()))
	}
	return ret
}
