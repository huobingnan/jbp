package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

// IntegerConstant 整型常量池常量
type IntegerConstant struct {
	value int
}

func (i *IntegerConstant) Tag() int {
	return ConstantIntegerInfo
}

func (i *IntegerConstant) Value() interface{} {
	return i.value
}

func (i *IntegerConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Integer_info: %d>", i.value)
}

func (i *IntegerConstant) GoString() string {
	return i.String()
}

func _newIntegerConstant(r *reader.ByteCodeReader) *IntegerConstant {
	if i, ok := r.ReadU4(); ok {
		return &IntegerConstant{int(i)}
	} else {
		panic("Read integer constant error")
	}
}
