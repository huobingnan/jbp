package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type StringConstant struct {
	index uint16
}

func (s *StringConstant) Tag() int {
	return ConstantStringInfo
}

func (s *StringConstant) Value() interface{} {
	return s.index
}

func (s *StringConstant) String() string {
	return fmt.Sprintf("<CONSTANT_String_info: @%d>", s.index)
}

func (s *StringConstant) GoString() string {
	return s.String()
}

func NewStringConstant(r *reader.ByteCodeReader) *StringConstant {
	if index, ok := r.ReadU2(); ok {
		return &StringConstant{
			index: index,
		}
	} else {
		panic("Read string constant error")
	}
}
