package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	StackMapTableEntries = "entries"
)

type StackMapFrame uint16

type StackMapTableAttribute struct {
	length  uint32
	entries []StackMapFrame
}

func (self *StackMapTableAttribute) Name() string { return StackMapTable }

func (self *StackMapTableAttribute) Length() uint32 { return 0 }

func (self *StackMapTableAttribute) Get(key string) interface{} {
	switch key {
	case StackMapTableEntries:
		return self.entries
	default:
		return nil
	}
}

func NewStackMapTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *StackMapTableAttribute {
	var ok bool
	ret := new(StackMapTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read stack map table attribute error",
			"can't read length info", r.Offset()))
	}
	numberOfEntries, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read stack map table attribute error",
			"can't read number_of_entries info", r.Offset()))
	}
	ret.entries = make([]StackMapFrame, 0, numberOfEntries)
	for i := uint16(0); i < numberOfEntries; i++ {
		offset, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read stack map table attribute error",
				"can't read offset info", r.Offset()))
		}
		ret.entries = append(ret.entries, StackMapFrame(offset))
	}
	return ret
}
