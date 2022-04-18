package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

type LongConstant struct {
	LongValue int64
}

func (l *LongConstant) Tag() int {
	return ConstantLongInfo
}

func (l *LongConstant) Value() interface{} {
	return l.LongValue
}

func (l *LongConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Long_info: %d>", l.LongValue)
}

func _newLongConstant(r *reader.ByteCodeReader) *LongConstant {
	if long, ok := r.ReadU8(); ok {
		return &LongConstant{
			LongValue: int64(long),
		}
	} else {
		panic("Read long constant error")
	}
}
