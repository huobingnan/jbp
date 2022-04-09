package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type LongConstant struct {
	value int64
}

func (l *LongConstant) Tag() int {
	return ConstantLongInfo
}

func (l *LongConstant) Value() interface{} {
	return l.value
}

func (l *LongConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Long_info: %d>", l.value)
}

func _newLongConstant(r *reader.ByteCodeReader) *LongConstant {
	if long, ok := r.ReadU8(); ok {
		return &LongConstant{
			value: int64(long),
		}
	} else {
		panic("Read long constant error")
	}
}
