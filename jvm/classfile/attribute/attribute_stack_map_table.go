package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

type StackMapFrame uint16

type StackMapTableAttribute struct {
	length  uint32
	Entries []StackMapFrame
}

func (s *StackMapTableAttribute) Name() string { return StackMapTable }

func (s *StackMapTableAttribute) Length() uint32 { return 0 }

func NewStackMapTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *StackMapTableAttribute {
	var ok bool
	ret := new(StackMapTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read stack map Table attribute error",
			"can't read Length info", r.Offset()))
	}
	numberOfEntries, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read stack map Table attribute error",
			"can't read number_of_entries info", r.Offset()))
	}
	if numberOfEntries >= 1 {
		ret.Entries = make([]StackMapFrame, 1, numberOfEntries)
	} else {
		ret.Entries = make([]StackMapFrame, 1)
	}
	// NOTICE: !!!!!!!!! numberOfEntries-1
	for i := uint16(0); i < numberOfEntries-1; i++ {
		offset, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read stack map Table attribute error",
				"can't read offset info", r.Offset()))
		}
		ret.Entries = append(ret.Entries, StackMapFrame(offset))
	}
	return ret
}
