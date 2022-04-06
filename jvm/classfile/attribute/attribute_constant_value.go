package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	ConstantValueIndex = "index"
)

type ConstantValueAttribute struct {
	length uint32
	index  uint16
}

func (self *ConstantValueAttribute) Name() string { return ConstantValue }

func (self *ConstantValueAttribute) Length() uint32 { return self.length }

func (self *ConstantValueAttribute) Get(key string) interface{} {
	switch key {
	case ConstantValueIndex:
		return self.index
	default:
		return nil
	}
}

func NewConstantValueAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *ConstantValueAttribute {
	var ok bool
	ret := new(ConstantValueAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read constant value attribute error", "can't read length info", r.Offset()))
	}
	ret.index, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read constant value attribute error", "can't read constantvalue_index info",
			r.Offset()))
	}
	return ret
}
