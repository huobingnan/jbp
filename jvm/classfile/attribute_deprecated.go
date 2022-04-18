package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
)

// DeprecatedAttribute 这是一个bool属性，只有存在于不存在，没有属性值
type DeprecatedAttribute struct {
}

func (d *DeprecatedAttribute) Name() string { return Deprecated }

func (d *DeprecatedAttribute) Length() uint32 { return 0 }

func (d *DeprecatedAttribute) Get(key string) interface{} { return nil }

func _newDeprecatedAttribute(r *reader.ByteCodeReader, cp ConstantPool) *DeprecatedAttribute {
	r.ReadU4()
	return new(DeprecatedAttribute)
}
