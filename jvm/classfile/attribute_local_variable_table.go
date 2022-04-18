package classfile

import (
	"github.com/huobingnan/jbp/jvm/reader"
)

type LocalVariableInfo struct {
	Length uint16
	// start_pc和length属性分别代表了这个局部变量的生命周期开始的字节码偏移量及其作用范围覆盖的长度
	// 两者结合起来就是这个局部变量在字节码之中的作用域范围
	StartPc uint16

	// name_index和descriptor_index都是指向常量池中CONSTANT_Utf8_info型常量的索引
	// 分别代表了局部变量的名称以及这个局部变量的描述符
	NameIndex       uint16
	DescriptorIndex uint16
	// index是这个局部变量在栈帧的局部变量表中变量槽的位置
	// 当这个变量数据类型是64位类型时 (double和long)，它占用的变量槽为index和index+1两个。
	Index uint16
}

type LocalVariableTableAttribute struct {
	length uint32
	Table  []LocalVariableInfo
}

func (l *LocalVariableTableAttribute) Name() string { return LocalVariableTable }

func (l *LocalVariableTableAttribute) Length() uint32 { return l.length }

func readLocalVariableInfo(r *reader.ByteCodeReader) LocalVariableInfo {
	info := LocalVariableInfo{}
	var ok bool
	info.StartPc, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read start_pc info", r.Offset()))
	}
	info.Length, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read Length info", r.Offset()))
	}
	info.NameIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read name_index info", r.Offset()))
	}
	info.DescriptorIndex, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error",
			"can't read descriptor_index info", r.Offset()))
	}
	info.Index, ok = r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read Index info", r.Offset()))
	}
	return info
}

func _newLocalVariableTableAttribute(r *reader.ByteCodeReader, cp ConstantPool) *LocalVariableTableAttribute {
	var ok bool
	ret := new(LocalVariableTableAttribute)
	ret.length, ok = r.ReadU4()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read Length info", r.Offset()))
	}
	tableLength, ok := r.ReadU2()
	if !ok {
		panic(ErrorMsgFmt("Read local variable info error", "can't read table_length info", r.Offset()))
	}
	ret.Table = make([]LocalVariableInfo, 0, tableLength)
	for i := uint16(0); i < tableLength; i++ {
		ret.Table = append(ret.Table, readLocalVariableInfo(r))
	}
	return ret
}
