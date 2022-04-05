package attribute

import (
	"bytecodeparser/jvm/classfile/constantpool"
	"bytecodeparser/jvm/classfile/reader"
)

const (
	LineNumberTableLineNumberTable = "line_number_table"
)

// LineNumberTableAttribute 用于描述Java源码行号与字节码行号(字节码的偏移量)之间的对应关系
type LineNumberTableAttribute struct {
	length          uint32
	lineNumberTable [][2]uint16 // 包含了start_pc和Java源程序line_number的对应，索引0存储start_pc，索引1存储line_number
}

func (l *LineNumberTableAttribute) Name() string { return LineNumberTable }

func (l *LineNumberTableAttribute) Length() uint32 { return l.length }

func (l *LineNumberTableAttribute) Get(key string) interface{} {
	switch key {
	case LineNumberTableLineNumberTable:
		return l.lineNumberTable
	default:
		return nil
	}
}

// NewLineNumberTableAttribute 从Class File中读取LineNumberTable的属性
func NewLineNumberTableAttribute(r *reader.ByteCodeReader, cp constantpool.ConstantPool) *LineNumberTableAttribute {
	var ok bool
	ret := new(LineNumberTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic("Read line number table attribute error (can't read length info)")
	}
	tableLength, ok := r.ReadU2()
	if !ok {
		panic("Read line number table attribute error (can't read table length info)")
	}
	ret.lineNumberTable = make([][2]uint16, 0, tableLength)
	for i := uint16(0); i < tableLength; i++ {
		startPc, ok := r.ReadU2()
		if !ok {
			panic("Read line number table attribute error (can't read start_pc info)")
		}
		lineNumber, ok := r.ReadU2()
		if !ok {
			panic("Read line number table attribute error (can't read line_number info)")
		}
		ret.lineNumberTable = append(ret.lineNumberTable, [2]uint16{startPc, lineNumber})
	}
	return ret
}