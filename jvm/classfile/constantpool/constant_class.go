package constantpool

import (
	"bytecodeparser/jvm/classfile/reader"
	"fmt"
)

// ClassConstant Java类常量池常量
type ClassConstant struct {
	Index uint16 // 这个类的符号引用，指向某个UTF8字符串常量在常量池中的索引值
}

func (c *ClassConstant) Tag() int {
	return ConstantClassInfo
}

func (c *ClassConstant) Value() interface{} {
	return c.Index
}

func (c *ClassConstant) String() string {
	return fmt.Sprintf("<CONSTANT_Class_info: @%d>", c.Index)
}

func (c *ClassConstant) GoString() string {
	return c.String()
}

func _newClassConstant(r *reader.ByteCodeReader) *ClassConstant {
	if index, ok := r.ReadU2(); ok {
		return &ClassConstant{
			Index: index,
		}
	} else {
		panic("Read class constant error")
	}
}
