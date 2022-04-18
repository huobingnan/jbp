package classfile

import (
	"bytecodeparser/jvm/reader"
)

// ExceptionTableAttribute 异常表属性
// 记录了Java程序运行时，异常的捕获范围，以及出现异常后的指令跳转信息
type ExceptionTableAttribute struct {
	StartPc   uint16 // 异常捕获的起始位置和结束位置（starPc的上一行)
	EndPc     uint16
	HandlerPc uint16 // 异常处理代码跳转的位置
	CatchType uint16 // 捕获异常的类型
}

func (e *ExceptionTableAttribute) Name() string { return ExceptionTable }

func (e *ExceptionTableAttribute) Length() uint32 { return uint32(4) }

func _newExceptionTableAttribute(r *reader.ByteCodeReader, cp ConstantPool) *ExceptionTableAttribute {
	var ok bool
	ret := new(ExceptionTableAttribute)
	ret.StartPc, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read exception Table error", "can't read start_pc info", r.Offset()))
	}
	ret.EndPc, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read exception Table error", "can't read end_pc info", r.Offset()))
	}
	ret.HandlerPc, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read exception Table error", "can't read handler_pc info", r.Offset()))
	}
	ret.CatchType, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read exception Table error", "can't read catch_type info", r.Offset()))
	}
	return ret
}
