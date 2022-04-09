package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type StringConstant struct {
	Index uint16
}

func (s *StringConstant) Tag() int {
	return ConstantStringInfo
}

func (s *StringConstant) Value() interface{} {
	return s.Index
}

func (s *StringConstant) String() string {
	return fmt.Sprintf("<CONSTANT_String_info: @%d>", s.Index)
}

func (s *StringConstant) GoString() string {
	return s.String()
}

func _newStringConstant(r *reader.ByteCodeReader) *StringConstant {
	if index, ok := r.ReadU2(); ok {
		return &StringConstant{
			Index: index,
		}
	} else {
		panic("Read string constant error")
	}
}
