package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	CodeMaxStack             = "max_stack"
	CodeMaxLocals            = "max_locals"
	CodeCodeLength           = "code_length"
	CodeCode                 = "code"
	CodeExceptionTableLength = "exception_table_length"
	CodeExceptionTable       = "exception_table"
	CodeAttributesCount      = "attributes_count"
	CodeAttributes           = "attributes"
)

// CodeAttribute JVM Code属性
type CodeAttribute struct {
	length         uint32
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []ExceptionTableAttribute
	attributes     []Attribute
}

func (c *CodeAttribute) Name() string { return Code }

func (c *CodeAttribute) Length() uint32 { return c.length }

func (c *CodeAttribute) Get(key string) interface{} {
	switch key {
	case CodeMaxStack:
		return c.maxStack
	case CodeMaxLocals:
		return c.maxLocals
	case CodeCodeLength:
		return len(c.code)
	case CodeCode:
		return c.code
	case CodeExceptionTableLength:
		return len(c.exceptionTable)
	case CodeExceptionTable:
		return c.exceptionTable
	case CodeAttributesCount:
		return len(c.attributes)
	case CodeAttributes:
		return c.attributes
	default:
		return nil
	}
}

func NewCodeAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *CodeAttribute {
	ret := new(CodeAttribute)
	if l, ok := r.ReadU4(); ok {
		ret.length = l
	} else {
		panic(ErrorMsgFmt("Read code attribute error", "can't read length info", r.Offset()))
	}
	if ms, ok := r.ReadU2(); ok {
		ret.maxStack = ms
	} else {
		panic(ErrorMsgFmt("Read code attribute error", "can't read max_stack info", r.Offset()))
	}
	if ml, ok := r.ReadU2(); ok {
		ret.maxLocals = ml
	} else {
		panic(ErrorMsgFmt("Read code attribute error", "can't read max_locals info", r.Offset()))
	}
	if cl, ok := r.ReadU4(); ok {
		if cl > 0 {
			// 读取代码
			// 这里进行了损失精度的类型转换，可能会出问题
			if code, ok := r.ReadAny(int(cl)); ok {
				ret.code = code
			} else {
				panic(ErrorMsgFmt("Read code attribute error", "can't read code info", r.Offset()))
			}
		}
	} else {
		panic(ErrorMsgFmt("Read code attribute error", "can't read code_length info", r.Offset()))
	}
	if el, ok := r.ReadU2(); ok {
		ret.exceptionTable = make([]ExceptionTableAttribute, el)
		for i := 0; i < int(el); i++ {
			ret.exceptionTable[i] = *NewExceptionTableAttribute(r, cp)
		}
	}
	// 读取其他属性
	if ac, ok := r.ReadU2(); ok {
		ret.attributes = make([]Attribute, ac)
		for i := 0; i < int(ac); i++ {
			ret.attributes[i] = New(r, cp)
		}
	}
	return ret
}
