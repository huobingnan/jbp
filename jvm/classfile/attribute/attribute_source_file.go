package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	SourceFileIndex = "index"
)

// SourceFileAttribute 源码名称属性
type SourceFileAttribute struct {
	length          uint32
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) Name() string { return SourceFile }

func (self *SourceFileAttribute) Length() uint32 { return self.length }

func (self *SourceFileAttribute) Get(key string) interface{} {
	switch key {
	case SourceFileIndex:
		return self.sourceFileIndex
	default:
		return nil
	}
}

func NewSourceFileAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *SourceFileAttribute {
	var ok bool
	ret := new(SourceFileAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read source file attribute error", "can't read length info", r.Offset()))
	}
	ret.sourceFileIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read source file attribute error",
			"can't read source_file_index info", r.Offset()))
	}
	return ret
}
