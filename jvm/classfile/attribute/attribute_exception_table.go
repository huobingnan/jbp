package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	ExceptionTableAttributeStartPc   = "start_pc"
	ExceptionTableAttributeEndPc     = "end_pc"
	ExceptionTableAttributeHandlerPc = "handler_pc"
	ExceptionTableAttributeCatchType = "catch_type"
)

// ExceptionTableAttribute 异常表属性
// 记录了Java程序运行时，异常的捕获范围，以及出现异常后的指令跳转信息
type ExceptionTableAttribute struct {
	startPc, endPc uint16 // 异常捕获的起始位置和结束位置（starPc的上一行)
	handlerPc      uint16 // 异常处理代码跳转的位置
	catchType      uint16 // 捕获异常的类型
}

func (self *ExceptionTableAttribute) Name() string { return ExceptionTable }

func (self *ExceptionTableAttribute) Length() uint32 { return uint32(4) }

func (self *ExceptionTableAttribute) Get(key string) interface{} {
	switch key {
	case ExceptionTableAttributeStartPc:
		return self.startPc
	case ExceptionTableAttributeEndPc:
		return self.endPc
	case ExceptionTableAttributeHandlerPc:
		return self.handlerPc
	case ExceptionTableAttributeCatchType:
		return self.catchType
	default:
		return nil
	}
}

func NewExceptionTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *ExceptionTableAttribute {
	var ok bool
	ret := new(ExceptionTableAttribute)
	ret.startPc, ok = r.ReadU2()
	if !ok {
		panic("Read exception table error (can't read start_pc)")
	}
	ret.endPc, ok = r.ReadU2()
	if !ok {
		panic("Read exception table error (can't read end_pc)")
	}
	ret.handlerPc, ok = r.ReadU2()
	if !ok {
		panic("Read exception table error (can't read handler_pc)")
	}
	ret.catchType, ok = r.ReadU2()
	if !ok {
		panic("Read exception table error (can't read catch_type)")
	}
	return ret
}
