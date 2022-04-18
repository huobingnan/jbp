package classfile

import (
	"bytecodeparser/jvm/reader"
)

// ExceptionsAttribute 标注在方法签名后的，该方法可能会抛出的异常
type ExceptionsAttribute struct {
	length     uint32   // 长度
	IndexTable []uint16 // 存储常量池中异常类型(CONSTANT_Class_info)的索引
}

func (self *ExceptionsAttribute) Name() string { return Exceptions }

func (self *ExceptionsAttribute) Length() uint32 { return self.length }

// _newExceptionsAttribute 从Class File中读取一个Exception
// NOTICE：读取时忽略属性的名称，直接从长度属性开始
func _newExceptionsAttribute(r *reader.ByteCodeReader, cp ConstantPool) *ExceptionsAttribute {
	length, ok := r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read exceptions attribute error", "can't read Length info", r.Offset()))
	}
	numberOfExceptions, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read exceptions attribute error", "can't read number_of_exceptions info",
			r.Offset()))
	}
	ret := new(ExceptionsAttribute)
	ret.length = length
	ret.IndexTable = make([]uint16, 0, numberOfExceptions)
	for i := uint16(0); i < numberOfExceptions; i++ {
		index, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read exceptions attribute error", "can't read exception_index info",
				r.Offset()))
		}
		ret.IndexTable = append(ret.IndexTable, index)
	}
	return ret
}
