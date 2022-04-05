package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

type StackMapTableAttribute struct {
}

func (self *StackMapTableAttribute) Name() string { return StackMapTable }

func (self *StackMapTableAttribute) Length() uint32 { return 0 }

func (self *StackMapTableAttribute) Get(key string) interface{} { return nil }

func NewStackMapTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *StackMapTableAttribute {
	// TODO
	panic("TODO")
}
