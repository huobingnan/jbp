package classfile

import (
	"github.com/huobingnan/jbp/jvm/reader"
)

// SyntheticAttribute 这是一个bool类型的属性，只有存在于不存在，没有属性值
type SyntheticAttribute struct {
}

func (s *SyntheticAttribute) Name() string { return Synthetic }

func (s *SyntheticAttribute) Length() uint32 { return 0 }

func (s *SyntheticAttribute) Get(key string) interface{} { return nil }

func _newSyntheticAttribute(r *reader.ByteCodeReader, cp ConstantPool) *SyntheticAttribute {
	r.ReadU4()
	return new(SyntheticAttribute)
}
