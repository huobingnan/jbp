package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

// CodeAttribute JVM Code属性
type CodeAttribute struct {
	length         uint32
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableAttribute
	Attributes     []JvmClassFileAttribute
}

func (c *CodeAttribute) Name() string { return Code }

func (c *CodeAttribute) Length() uint32 { return c.length }

func _newCodeAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *CodeAttribute {
	ret := new(CodeAttribute)
	if l, ok := r.ReadU4(); ok {
		ret.length = l
	} else {
		panic(ErrorMsgFmt("Read Code attribute error", "can't read Length info", r.Offset()))
	}
	if ms, ok := r.ReadU2(); ok {
		ret.MaxStack = ms
	} else {
		panic(ErrorMsgFmt("Read Code attribute error", "can't read max_stack info", r.Offset()))
	}
	if ml, ok := r.ReadU2(); ok {
		ret.MaxLocals = ml
	} else {
		panic(ErrorMsgFmt("Read Code attribute error", "can't read max_locals info", r.Offset()))
	}
	if cl, ok := r.ReadU4(); ok {
		if cl > 0 {
			// 读取代码
			// 这里进行了损失精度的类型转换，可能会出问题
			if code, ok := r.ReadAny(int(cl)); ok {
				ret.Code = code
			} else {
				panic(ErrorMsgFmt("Read Code attribute error", "can't read Code info", r.Offset()))
			}
		}
	} else {
		panic(ErrorMsgFmt("Read Code attribute error", "can't read code_length info", r.Offset()))
	}
	if el, ok := r.ReadU2(); ok {
		ret.ExceptionTable = make([]ExceptionTableAttribute, el)
		for i := 0; i < int(el); i++ {
			ret.ExceptionTable[i] = *_newExceptionTableAttribute(r, cp)
		}
	}
	// 读取其他属性
	if ac, ok := r.ReadU2(); ok {
		ret.Attributes = make([]JvmClassFileAttribute, ac)
		for i := 0; i < int(ac); i++ {
			ret.Attributes[i] = New(r, cp)
		}
	}
	return ret
}
