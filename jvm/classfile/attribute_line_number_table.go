package classfile

import (
	"bytecodeparser/jvm/classfile/reader"
)

// LineNumberInfo 包含了start_pc和Java源程序line_number的对应
// 索引0存储start_pc，索引1存储line_number
type LineNumberInfo [2]uint16

func (self *LineNumberInfo) ByteCodeLineNumber() uint16 { return (*self)[0] }

func (self *LineNumberInfo) SourceFileLineNumber() uint16 { return (*self)[1] }

// LineNumberTableAttribute 用于描述Java源码行号与字节码行号(字节码的偏移量)之间的对应关系
type LineNumberTableAttribute struct {
	length uint32
	Table  []LineNumberInfo
}

func (l *LineNumberTableAttribute) Name() string { return LineNumberTable }

func (l *LineNumberTableAttribute) Length() uint32 { return l.length }

// _newLineNumberTableAttribute 从Class File中读取LineNumberTable的属性
func _newLineNumberTableAttribute(r *reader.ByteCodeReader, cp ConstantPool) *LineNumberTableAttribute {
	var ok bool
	ret := new(LineNumberTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read line number Table attribute error",
			"can't read Length info", r.Offset()))
	}
	tableLength, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read line number Table attribute error",
			"can't read table_length info", r.Offset()))
	}
	ret.Table = make([]LineNumberInfo, 0, tableLength)
	for i := uint16(0); i < tableLength; i++ {
		startPc, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read line number Table attribute error",
				"can't read  start_cp info", r.Offset()))
		}
		lineNumber, ok := r.ReadU2()
		if !ok {
			panic(ErrorMsgFmt("Read line number Table attribute error",
				"can't read line_number info", r.Offset()))
		}
		ret.Table = append(ret.Table, LineNumberInfo([2]uint16{startPc, lineNumber}))
	}
	return ret
}
